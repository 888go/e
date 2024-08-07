package e

import (
	"fmt"
	"time"
)

type E时间统计类 struct {
	t time.Time
}

func X创建时间统计类() *E时间统计类 {
	t := new(E时间统计类)
	t.X开始()
	return t
}
func (this *E时间统计类) X开始() {
	this.t = time.Now()
}

func (this *E时间统计类) X取毫秒() string {
	return fmt.Sprintf("%d", time.Since(this.t).Milliseconds())
}
func (this *E时间统计类) X取秒() string {
	return fmt.Sprintf("%.3f", float64(time.Since(this.t).Milliseconds())/float64(1000))
}
