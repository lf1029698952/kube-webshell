package controllers

import (
	"bytes"
	inCon "context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"io"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/kubernetes/pkg/util/interrupt"
	"net/http"
	"strings"
)

const TOKEN_HEADER = "token"

func (self TerminalSockjs) Read(p []byte) (int, error) {
	var reply string
	var msg map[string]uint16
	reply, err := self.conn.Recv()
	if err != nil {
		return 0, err
	}
	if err := json.Unmarshal([]byte(reply), &msg); err != nil {
		return copy(p, reply), nil
	} else {
		self.sizeChan <- &remotecommand.TerminalSize{
			msg["cols"],
			msg["rows"],
		}
		return 0, nil
	}
}

func (self TerminalSockjs) Write(p []byte) (int, error) {
	var err error
	if strings.Contains(string(p), "OCI runtime exec failed: exec failed: container_linux.go:344") {
		beego.Info("/bin/bash not support")
	} else {
		err = self.conn.Send(string(p))
	}
	return len(p), err
}

type TerminalSockjs struct {
	conn      sockjs.Session
	sizeChan  chan *remotecommand.TerminalSize
	context   string
	namespace string
	pod       string
	container string
}

// 实现tty size queue
func (self *TerminalSockjs) Next() *remotecommand.TerminalSize {
	size := <-self.sizeChan
	beego.Debug(fmt.Sprintf("terminal size to width: %d height: %d", size.Width, size.Height))
	return size
}

func buildConfigFromContextFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}

// 处理输入输出与sockjs 交互
func Handler(t *TerminalSockjs, cmd string) error {
	config, err := buildConfigFromContextFlags(t.context, beego.AppConfig.String("kubeconfig"))
	clientset, err := kubernetes.NewForConfig(config)
	pods, podQerr := clientset.CoreV1().Pods(t.namespace).Get(inCon.TODO(), t.pod, metav1.GetOptions{})
	if podQerr != nil || pods.Size() <= 0 {
		logs.Info("pod not found", t.pod)
		return errors.New("pod: " + t.pod + " not found")
	}
	if err != nil {
		return err
	}
	groupversion := schema.GroupVersion{
		Group:   "",
		Version: "v1",
	}
	config.GroupVersion = &groupversion
	config.APIPath = "/api"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = scheme.Codecs
	/*config.NegotiatedSerializer = serializer.CodecFactory{: scheme.Codecs}*/
	restclient, err := rest.RESTClientFor(config)
	if err != nil {
		return err
	}
	fn := func() error {
		req := restclient.Post().
			Resource("pods").
			Name(t.pod).
			Namespace(t.namespace).
			SubResource("exec").
			Param("container", t.container).
			Param("stdin", "true").
			Param("stdout", "true").
			Param("stderr", "true").
			Param("command", cmd).Param("tty", "true")
		req.VersionedParams(
			&v1.PodExecOptions{
				Container: t.container,
				Command:   []string{},
				Stdin:     true,
				Stdout:    true,
				Stderr:    true,
				TTY:       true,
			},
			scheme.ParameterCodec,
		)
		executor, err := remotecommand.NewSPDYExecutor(
			config, http.MethodPost, req.URL(),
		)
		if err != nil {
			return err
		}
		return executor.Stream(remotecommand.StreamOptions{
			Stdin:             t,
			Stdout:            t,
			Stderr:            t,
			Tty:               true,
			TerminalSizeQueue: t,
		})
	}
	//i m not sure why restore?
	/*inFd, _ := term.GetFdInfo(t.conn)
	state, err := term.SaveState(inFd)*/

	return interrupt.Chain(nil, func() {
		/*term.RestoreTerminal(inFd, state)*/
	}).Run(fn)
}

// 实现http.handler 接口获取入参
func (self TerminalSockjs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := r.FormValue("context")
	namespace := r.FormValue("namespace")
	pod := r.FormValue("pod")
	container := r.FormValue("container")
	Sockjshandler := func(session sockjs.Session) {
		defer session.Close(0, "exit close 0, pod not found")

		t := &TerminalSockjs{session, make(chan *remotecommand.TerminalSize),
			context, namespace, pod, container}
		if err := Handler(t, "/bin/bash"); err != nil {
			beego.Error(err)
			beego.Error(Handler(t, "/bin/sh"))
		}
	}

	sockjs.NewHandler("/terminal/ws", sockjs.DefaultOptions, Sockjshandler).ServeHTTP(w, r)
}

//终端过滤器 校验token 是否正确
func FilterToken(ctx *context.Context) {
	beego.Info("tiger==========>", ctx.Request.URL.RequestURI())
	if len(beego.AppConfig.String("gateway-address")) <= 0 {
		beego.Info("congratulations, there are no auth")
		return
	}
	if strings.Contains(ctx.Request.URL.RequestURI(), TOKEN_HEADER) {
		token := getTokenByUri(ctx.Request.URL.RequestURI())
		if len(token) <= 0 {
			logs.Warning("not contain token")
			http.Error(ctx.ResponseWriter, "auth failed", http.StatusUnauthorized)
		}
		url := "http://" + beego.AppConfig.String("gateway-address") + ":" + beego.AppConfig.String("gateway-port") + "/api/oauth/validate-token?token=" + token
		resp, err := http.Get(url)
		if err != nil {
			logs.Error("can not get connection to gateway")
			logs.Error(err)
			http.Error(ctx.ResponseWriter, "auth failed", http.StatusUnauthorized)
		}
		var buffer bytes.Buffer
		for {
			content := make([]byte, 1024)
			_, err := resp.Body.Read(content)
			buffer.Write(content)
			if err == io.EOF {
				break
			}
		}
		logs.Info(buffer.String())
		if resp.StatusCode != http.StatusOK {
			logs.Error("invalid token, check please!!")
			http.Error(ctx.ResponseWriter, "auth failed", http.StatusUnauthorized)
		}
	} else {
		beego.Warning("not contain token")
		http.Error(ctx.ResponseWriter, "auth failed", http.StatusUnauthorized)
	}

}

func getTokenByUri(uri string) string {
	buffer := strings.Split(uri, "&")
	index := -1
	for i, value := range buffer {
		if strings.Contains(value, TOKEN_HEADER) {
			index = i
			break
		}
	}
	if index >= 0 {
		token := strings.Split(buffer[index], "=")
		return token[1]
	}
	return ""
}
