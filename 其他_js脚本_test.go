package e

import (
	"testing"
)

func TestRunJs(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := X运行JS脚本("function get(p){return p}", "get", "aaaaaaa"+X到文本(i))
		X调试输出(i, str)
	}
}
