package e

import (
	"testing"
)

func TestNew文本型数组(t *testing.T) {
	strarr := X创建文本数组()
	X调试输出(strarr)
	strarr.X加入成员("3")
	strarr.X加入成员("2")
	strarr.X加入成员("5")
	strarr.X加入成员("1")
	strarr.X加入成员("4")
	strarr.X插入成员(1, "6")

	X调试输出(strarr)
	X调试输出(strarr.X取数组成员数())
	strarr.X删除成员(1)
	X调试输出(strarr.X取数组成员数())

	//strarr.X清除数组()
	strarr.X数组排序(true)
	X调试输出(strarr)
	strarr.X数组排序(false)
	X调试输出(strarr)
	X调试输出(strarr.X取值(0))

	for i := 0; i < strarr.X取数组成员数(); i++ {
		X调试输出(i, strarr.X取值(i))
		X调试输出(i, strarr.X取值(i))
	}
}
