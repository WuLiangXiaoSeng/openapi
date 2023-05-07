package main

import (
	"wuliangxiaoseng/openapi/internal/api"
	"wuliangxiaoseng/openapi/internal/common"
)

func main() {
	common.Println("Hello world")
	common.Debug()

	app := api.NewServerApp()

	app.Start(true)
}
