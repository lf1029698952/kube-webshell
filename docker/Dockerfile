FROM alpine:latest

LABEL maintainer="kube webshell by lifenga@zbj.com"

RUN apk add -U tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime

COPY kube-webshell-linux-amd64 /opt/kube/kube-webshell-linux-amd64
COPY static /opt/kube/static
COPY views /opt/kube/views

WORKDIR /opt/kube
CMD ./kube-webshell-linux-amd64