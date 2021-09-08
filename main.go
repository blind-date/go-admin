package main

import (
	"go-admin/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6

// 本地启动命令：go run main.go server -c=config/settings.yml
// 本地初始化命令：go run main.go migrate -c=config/settings.yml
func main() {
	cmd.Execute()
}
