package routes

import "github.com/kataras/iris"

func Web(app *iris.Application)  {
	app.Get("/", func(ctx iris.Context) {
		_, _ = ctx.HTML("<h1>首页</h1>")
	})
}
