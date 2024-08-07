package e

//定时任务
//提供周期性任务执行功能。使用方式类似于Linux下的Crontab或者Java里的Quartz。本对象线程安全。
import (
	"github.com/888go/e/internal/cable"
	"time"
)

type X时钟 struct {
}

func X创建时钟类() *X时钟 {
	this := new(X时钟)
	return this
}

func (this *X时钟) X创建(fn func() bool, 时钟周期 int64) {
	cable.SetInterval(fn, time.Duration(时钟周期)*time.Millisecond)
}

func (this *X时钟) X创建执行一次(fn func(), 时钟周期 int64) {
	cable.SetTimeout(fn, time.Duration(时钟周期)*time.Millisecond)
}

// 返回一个函数，无论调用多少次，它只会在指定的间隔内执行一次
func (this *X时钟) X时钟周期函数(fn func(), 时钟周期 int64) func() {
	return cable.Throttle(fn, time.Duration(时钟周期)*time.Millisecond, cable.ThrottleOptions{})
}

func X时钟_创建(fn func() bool, 时钟周期 int64) {
	cable.SetInterval(fn, time.Duration(时钟周期)*time.Millisecond)
}

func X时钟_创建执行一次(fn func(), 时钟周期 int64) {
	cable.SetTimeout(fn, time.Duration(时钟周期)*time.Millisecond)
}

// 返回一个函数，无论调用多少次，它只会在指定的间隔内执行一次
func X时钟_创建周期函数(fn func(), 时钟周期 int64) func() {
	return cable.Throttle(fn, time.Duration(时钟周期)*time.Millisecond, cable.ThrottleOptions{})
}
