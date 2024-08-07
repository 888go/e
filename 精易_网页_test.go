package e

import (
	"github.com/888go/goframe/frame/g"
	"testing"
)

func TestGclient(t *testing.T) {
	附加协议 := g.MapStrStr{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0"}
	测试2 := X网页访问("https://goframe.org/display/gf/HTTPClient", 0, "", nil, nil, 附加协议, "", 100)
	X调试输出(测试2)
	//X写到文件("C:/Users/admin/Desktop/测试.html", []byte(测试2))
}
