// Package ejs js引擎 用于运行js代码
package e

import (
	"github.com/robertkrimen/otto"
)

// 使用给定的this和参数调用给定的JavaScript。
func X运行JS脚本(脚本 string, 函数名称 string, 参数列表 ...interface{}) string {
	var err error
	vm := otto.New()
	_, err = vm.Run(脚本)
	if err != nil {
		panic(err)
	}
	value, err := vm.Call(函数名称, nil, 参数列表...)
	if err != nil {
		panic(err)
	}
	return value.String()
}
