package e

import (
	"fmt"
	"testing"
)

func TestNew配置项(t *testing.T) {
	var config X配置项
	text2 := `
[配置]
服务器=127.0.0.1
端口=9001
[redis]
addr=127.0.0.1
port=6379
`
	config.X加载配置项从内存(text2)
	所有节名 := config.X取节名()
	for i, 节名 := range 所有节名 {
		fmt.Println(i, 节名)
		所有项名 := config.X取项名(节名)
		for i, 项名 := range 所有项名 {
			fmt.Println(i, 节名, 项名, config.X读配置项(节名, 项名, ""))
		}
	}
}
