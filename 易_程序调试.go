package e

import (
	"fmt"
	"github.com/kr/pretty"
)

func X调试输出(a ...interface{}) {
	pretty.Println(a...)
}

func X调试格式化输出(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}
