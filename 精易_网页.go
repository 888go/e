package e

import (
	"github.com/888go/goframe/frame/g"
	网页类 "github.com/888go/goframe/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

// 网页访问
// 网址：完整的网页地址,必须包含http://或者https://
// 访问方式：0=GET 1=POST 2=HEAD 3=PUT  4=OPTIONS  5=DELETE  6=TRACE  7=CONNECT
// 提交信息：可以是map类型,或者文本类型,或者json文本
// 提交Cookies：设置提交时的cookie
// 返回Cookies：返回Cookies
// 附加协议头：一行一个请用换行符隔开,建议填写常量值或文本值,防止因传参引发错误
// 返回协议头：返回协议头
// 字节集提交：提交字节集数据
// 代理地址：代理地址，格式为:ip:port,例如:8.8.8.8:88
// 超时_秒：超时时间，单位秒，默认10秒
func X网页访问(网址 string, 访问方式 int, 提交信息 interface{}, 提交Cookies map[string]string, 返回Cookies *map[string]string, 附加协议头 map[string]string, 代理地址 string, 超时_秒 time.Duration) string {
	ctx := gctx.New()
	超时时长 := 超时_秒 * time.Second
	if 超时_秒 == 0 {
		超时时长 = 5 * time.Second
	}
	网页对象 := g.X网页类().X设置CookieMap(提交Cookies).X设置Map协议头(附加协议头).X代理(代理地址).X超时(超时时长)

	var 响应对象 *网页类.Response
	if 访问方式 == 0 {
		响应对象, _ = 网页对象.Get响应对象(ctx, 网址, 提交信息)
	}
	if 访问方式 == 1 {
		响应对象, _ = 网页对象.Post响应对象(ctx, 网址, 提交信息)
	}
	if 访问方式 == 2 {
		响应对象, _ = 网页对象.Head响应对象(ctx, 网址, 提交信息)
	}
	if 访问方式 == 3 {
		响应对象, _ = 网页对象.Put响应对象(ctx, 网址, 提交信息)
	}
	if 访问方式 == 4 {
		响应对象, _ = 网页对象.Options响应对象(ctx, 网址, 提交信息)
	}
	if 访问方式 == 5 {
		响应对象, _ = 网页对象.Delete响应对象(ctx, 网址, 提交信息)
	}
	if 访问方式 == 6 {
		响应对象, _ = 网页对象.Trace响应对象(ctx, 网址, 提交信息)
	}
	if 访问方式 == 7 {
		响应对象, _ = 网页对象.Connect响应对象(ctx, 网址, 提交信息)
	}
	if 返回Cookies != nil {
		*返回Cookies = 响应对象.X取CookieMap()
	}
	返回文本 := 响应对象.X取响应文本()
	响应对象.X关闭()
	return 返回文本
}
