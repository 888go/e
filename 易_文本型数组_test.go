package e

import (
	"testing"
)

func TestNew文本型数组(t *testing.T) {
	strarr := X创建文本数组()
	E调试输出(strarr)
	strarr.E加入成员("3")
	strarr.E加入成员("2")
	strarr.E加入成员("5")
	strarr.E加入成员("1")
	strarr.E加入成员("4")
	strarr.E插入成员(1, "6")

	E调试输出(strarr)
	E调试输出(strarr.E取数组成员数())
	strarr.E删除成员(1)
	E调试输出(strarr.E取数组成员数())

	//strarr.E清除数组()
	strarr.E数组排序(true)
	E调试输出(strarr)
	strarr.E数组排序(false)
	E调试输出(strarr)
	E调试输出(strarr.E取值(0))

	for i := 0; i < strarr.E取数组成员数(); i++ {
		E调试输出(i, strarr.E取值(i))
		E调试输出(i, strarr.E取值(i))
	}
}
