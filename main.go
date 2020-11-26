package main

import (
	"github.com/astaxie/beego"
	"kube-webshell/controllers"
	_ "kube-webshell/routers"
)

func main() {
	beego.InsertFilter("/terminal/*", beego.BeforeExec, controllers.FilterToken)
	beego.Run()
}
