module kube-webshell

go 1.15

require (
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/vcs v1.13.1 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/astaxie/beego v1.12.3
	github.com/golang/dep v0.5.4 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jmank88/nuts v0.4.0 // indirect
	github.com/nightlyone/lockfile v1.0.0 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/sdboyer/constext v0.0.0-20170321163424-836a14457353 // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	golang.org/x/sys v0.0.0-20201126233918-771906719818 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/igm/sockjs-go.v2 v2.1.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/kubernetes v0.19.0
)

replace (
	k8s.io/api => k8s.io/api v0.19.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.0
	k8s.io/apiserver => k8s.io/apiserver v0.19.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.0
	k8s.io/client-go => k8s.io/client-go v0.19.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.0
	k8s.io/code-generator => k8s.io/code-generator v0.19.0
	k8s.io/component-base => k8s.io/component-base v0.19.0
	k8s.io/cri-api => k8s.io/cri-api v0.19.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.0
	k8s.io/kubectl => k8s.io/kubectl v0.19.0
	k8s.io/kubelet => k8s.io/kubelet v0.19.0
	k8s.io/kubernetes => github.com/kubernetes/kubernetes v1.19.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.0
	k8s.io/metrics => k8s.io/metrics v0.19.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.0
)
