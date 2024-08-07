package ecore

import (
	"fmt"
	"github.com/kr/pretty"
)

func E调试输出(a ...interface{}) {
	pretty.Println(a...)
}

func E调试格式化输出(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}
