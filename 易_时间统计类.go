package e

import (
	"fmt"
	"time"
)

type X时间统计类 struct {
	t time.Time
}

func X创建时间统计类() *X时间统计类 {
	t := new(X时间统计类)
	t.X开始()
	return t
}
func (this *X时间统计类) X开始() {
	this.t = time.Now()
}

func (this *X时间统计类) X取毫秒() string {
	return fmt.Sprintf("%d", time.Since(this.t).Milliseconds())
}
func (this *X时间统计类) X取秒() string {
	return fmt.Sprintf("%.3f", float64(time.Since(this.t).Milliseconds())/float64(1000))
}
