package e

import (
	"fmt"
	"testing"
)

func Test_取现行时间(t *testing.T) {
	time := X取现行时间()
	t.Log("取年份", time.X取年份())
	t.Log("X取月份", time.X取月份())
	t.Log("X取日", time.X取日())
	t.Log("X取星期几", time.X取星期几())
	t.Log("X取小时", time.X取小时())
	t.Log("X取分钟", time.X取分钟())
	t.Log("X取秒", time.X取秒())
	t.Log("x取月天数", time.x取月天数())
	t.Log("现行时间", time.X时间到文本("Y-m-d H:i:s"))
	time = X到时间("2019-02-04 22:53:02")
	t.Log("取年份", time.X取年份())
	t.Log("X取月份", time.X取月份())
	t.Log("X取日", time.X取日())
	t.Log("X取星期几", time.X取星期几())
	t.Log("X取小时", time.X取小时())
	t.Log("X取分钟", time.X取分钟())
	t.Log("X取秒", time.X取秒())
	t.Log("x取月天数", time.x取月天数())
	t.Log("现行时间", time.X时间到文本("Y-m-d H:i:s"))

	time.X增减日期(1, 0, 0)
	t.Log("X增减时间", time.X时间到文本("Y-m-d H:i:s"))

	time.X增减时间(1, 1, 1)
	t.Log("X增减时间", time.X时间到文本("Y-m-d H:i:s"))

	time.X增减时间(-1, -1, -1)
	t.Log("X增减时间", time.X时间到文本("Y-m-d H:i:s"))

	time2 := X到时间("2022-03-04 22:53:02")
	t.Log("X大于", time.X大于(time2))
	t.Log("X小于", time.X小于(time2))

	time3 := X到时间("X增减时间 2020-03-04 22:53:02")
	t.Log("X等于", time.X等于(time3))

	time4 := X取现行时间()
	t.Log("X取时间戳", time4.X取时间戳())
	t.Log("X取时间戳毫秒", time4.X取时间戳毫秒())
	t.Log("X取时间戳微秒", time4.X取时间戳微秒())
	t.Log("X取时间戳纳秒", time4.X取时间戳纳秒())

	//t.Log("X到友好时间", time4.X到友好时间())

	t.Log("X到时间从时间戳", X到时间从时间戳(X取现行时间().X取时间戳()).X时间到文本(""))

	t.Log("X取时间戳", X取现行时间().X取时间戳())

	//time2 = X到时间("2019-02-04 22:53:02")
	//t.Log("X到友好时间", time2.X到友好时间())
	time2 = X取现行时间().X增减日期(0, 0, -10)
	t.Log("X到友好时间", time2.X到友好时间(), time2.X时间到文本("Y-m-d H:i:s"))
}

func ExampleE到友好时间() {
	time := X到时间("2019-02-04 22:53:02")

	t1 := time.X到友好时间()
	fmt.Println(t1)
}
