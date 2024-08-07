package e

import (
	"github.com/888go/goframe/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestGclient(t *testing.T) {
	ctx := gctx.New()
	X调试输出(g.X网页类().Get文本(ctx, "https://goframe.org/pages/viewpage.action?pageId=1114381"))
	X调试输出(X网页访问("http://www.baidu.com/", 0, "", nil, nil, nil, "", 0))
}
