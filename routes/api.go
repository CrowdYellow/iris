package routes

import "github.com/kataras/iris"

func Api(app *iris.Application)  {
	var main = corsSetting(app)
	main.Get("/", func(ctx iris.Context) {
		_, _ = ctx.HTML("<h1>首页</h1>")
	})
}
