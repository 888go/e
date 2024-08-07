package e

import (
	"github.com/888go/goframe/container/garray"
)

type X文本型数组 struct {
	X父类 *切片类.StrArray
}

type X数组 struct {
	X父类 *切片类.Array
}

type X整数型数组 struct {
	X父类 *切片类.IntArray
}

func X创建文本数组(并发安全 ...bool) *X文本型数组 {
	return &X文本型数组{X父类: 切片类.X创建文本(并发安全...)}
}
func (a *X文本型数组) X加入成员(s string) {
	a.X父类.Append别名(s)
}
func (a *X文本型数组) X取数组成员数() int {
	return a.X父类.X取长度()
}
func (a *X文本型数组) X删除成员(索引 int) {
	a.X父类.X删除(索引)
}
func (a *X文本型数组) X清除数组() {
	a.X父类.X清空()
}

func (a *X文本型数组) X数组排序(降序 bool) {
	a.X父类.X排序递增(降序)
}
func (a *X文本型数组) X插入成员(索引 int, 值 ...string) {
	a.X父类.X插入前面(索引, 值...)
}
func (a *X文本型数组) X取值(索引 int) string {
	返回值, _ := a.X父类.X取值2(索引)
	return 返回值
}

func X创建数组(并发安全 ...bool) *X数组 {
	return &X数组{X父类: 切片类.X创建(并发安全...)}
}
func (a *X数组) X加入成员(s interface{}) {
	a.X父类.Append别名(s)
}
func (a *X数组) X取数组成员数() int {
	return a.X父类.X取长度()
}
func (a *X数组) X删除成员(索引 int) {
	a.X父类.X删除(索引)
}
func (a *X数组) X清除数组() {
	a.X父类.X清空()
}
func (a *X数组) X插入成员(索引 int, 值 ...interface{}) {
	a.X父类.X插入前面(索引, 值...)
}
func (a *X数组) X取值(索引 int) interface{} {
	返回值, _ := a.X父类.X取值2(索引)
	return 返回值
}

func X创建整数数组(并发安全 ...bool) *X整数型数组 {
	return &X整数型数组{X父类: 切片类.X创建整数(并发安全...)}
}
func (a *X整数型数组) X加入成员(s int) {
	a.X父类.Append别名(s)
}
func (a *X整数型数组) X取数组成员数() int {
	return a.X父类.X取长度()
}
func (a *X整数型数组) X删除成员(索引 int) {
	a.X父类.X删除(索引)
}
func (a *X整数型数组) X清除数组() {
	a.X父类.X清空()
}

func (a *X整数型数组) X数组排序(降序 bool) {
	a.X父类.X排序递增(降序)
}
func (a *X整数型数组) X插入成员(索引 int, 值 ...int) {
	a.X父类.X插入前面(索引, 值...)
}
func (a *X整数型数组) X取值(索引 int) int {
	返回值, _ := a.X父类.X取值2(索引)
	return 返回值
}
