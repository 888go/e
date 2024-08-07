package e

import (
	"testing"
)

func TestNew时钟(t *testing.T) {
	return
	时钟 := New时钟()
	时间统计 := X创建时间统计类()
	i := 0
	时钟.X创建(func() bool {
		i++
		X延时(1000)
		if i > 10 {
			X调试输出("10次停止了")
			return false
		}
		X调试输出(i, 时间统计.X取毫秒())
		return true
	}, 100)

	时钟.X创建执行一次(func() {
		X调试输出("执行一次", 时间统计.X取毫秒())
	}, 100)

	时钟周期函数 := 时钟.X时钟周期函数(func() {
		X调试输出("3秒调用多少次都只执行1次", 时间统计.X取毫秒())
	}, 1000*3)

	时钟.X创建(func() bool {
		时钟周期函数()

		//X调试输出("调用时钟周期函数",时间统计.X取毫秒())
		return true
	}, 100)

	select {}
}
