package e

import (
	"testing"
)

func TestNew正则表达式(t *testing.T) {
	str := `aaa111bbb444ccc
	aaa222bbb555ccc
	aaa333bbb666ccc`

	var zz X正则表达式
	if zz.X创建(`aaaaa(.*?)bbb(.*?)ccc`, str) {
		X调试输出(zz.X取匹配数量())
		for i := 0; i < zz.X取匹配数量(); i++ {
			X调试输出(i, zz.X取子匹配文本(i, 1))
			X调试输出(i, zz.X取子匹配文本(i, 2))
			X调试输出(zz.X取子匹配文本(i, 3))
		}
		X调试输出(zz.X取子匹配文本(4, 3))
	}

	zz2, flag := X创建正则表达式(`aaa(.*?)bbb(.*?)ccc`, str)
	X调试输出(flag)
	X调试输出(zz2.X取匹配数量())
	for i := 0; i < zz2.X取匹配数量(); i++ {
		X调试输出(i, zz2.X取子匹配文本(i, 1))
		X调试输出(i, zz2.X取子匹配文本(i, 2))
		X调试输出(zz2.X取子匹配文本(i, 3))
	}
	X调试输出(zz2.X取子匹配文本(4, 3))

	arr := zz2.GetResult()

	X调试输出("GetResult", arr)

	zz3, flag := X创建正则表达式(`aaa`, str)

	arr2 := zz3.X替换("666")

	X调试输出("X替换", arr2)

}
