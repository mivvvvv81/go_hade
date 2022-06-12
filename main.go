package main

import (
	"project/app/console"
	"project/app/http"
	"project/framework"
	"project/framework/provider/app"
	"project/framework/provider/kernel"
)

func main() {
	container := framework.NewHadeContainer()
	container.Bind(&app.HadeAppProvider{})
	if engine, err := http.NewHttpEngine(); err != nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}
	console.RunCommand(container)
}
