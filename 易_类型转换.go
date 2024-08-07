package e

import (
	"github.com/gogf/gf/v2/util/gconv"
)

func X到字节集(value interface{}) []byte {
	return gconv.Bytes(value)
}
func X到字节(value interface{}) byte {
	return gconv.Byte(value)
}
func X到整数(value interface{}) int64 {
	return gconv.Int64(value)
}

func X到数值(value interface{}) float64 {
	return gconv.Float64(value)
}
func X到文本(value interface{}) string {
	return gconv.String(value)
}
func X到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error {
	return gconv.Struct(待转换的参数, 结构体指针)
}
