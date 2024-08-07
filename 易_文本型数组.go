package ecore

import "github.com/gogf/gf/v2/container/garray"

type StrArray struct {
	X父类 *garray.StrArray
}

func X创建文本数组(并发安全 ...bool) *StrArray {
	return &StrArray{X父类: garray.NewStrArray(并发安全...)}
}
func (a *StrArray) E加入成员(s string) {
	a.X父类.Append(s)
}
func (a *StrArray) E取数组成员数() int {
	return a.X父类.Len()
}

func (a *StrArray) E删除成员(欲删除的位置 int) {
	a.X父类.Remove(欲删除的位置)
}

func (a *StrArray) E清除数组() {
	a.X父类.Clear()
}

// E数组排序
// 对指定数值数组变量内的所有数组成员进行快速排序，不影响数组的维定义信息，排序结果存放回该数组变量。本命令为初级命令。
// 参数<1>的名称为“排序方向是否为从小到大”，类型为“逻辑型（bool）”，可以被省略。如果参数值为真，排序方向为从小到大，否则为从大到小。如果本参数被省略，默认值为真。
func (a *StrArray) E数组排序(排序方向是否为从小到大 bool) {
	a.X父类.Sort(排序方向是否为从小到大)
}

// E插入成员
// 参数<1>的名称为“欲插入的位置”，类型为“整数型（int）”。位置值从 1 开始，如果小于 1 或大于第一个参数所指定数组变量的成员数目 + 1，将不会插入任何数据。
// 参数<2>的名称为“欲插入的成员数据”，类型为“通用型（all）”，提供参数数据时可以同时提供数组或非数组数据。参数值所指定成员数据的数据类型必须能够与第一个参数所指定的数组变量相匹配。
func (a *StrArray) E插入成员(欲插入的位置 int, 欲插入的成员数据 string) {
	a.X父类.InsertBefore(欲插入的位置, 欲插入的成员数据)
}
func (a *StrArray) E取值(i int) string {
	返回值, _ := a.X父类.Get(i)
	return 返回值
}
