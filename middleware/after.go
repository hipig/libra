package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/lhlyu/libra/logger"
)

// 可以做统一响应处理 比如打印日志记录响应结果...
func After() context.Handler {
	return func(ctx iris.Context) {
		body := ctx.Recorder().Body() // 获取响应返回的内容
		logger.Log(ctx).WithField("result", string(body)).Debugln()
		ctx.Recorder().ResetBody() // 将响应体body内容置空
		ctx.Write(body)            // 重写写入body
	}
}
