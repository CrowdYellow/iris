package routes

import (
	"github.com/kataras/iris"
	"iris/app/Controllers"
)

func Api(app *iris.Application) {
	var main = corsSetting(app, "/api/")
	main.Get("/", func(ctx iris.Context) {
		_, _ = ctx.HTML("<h1>首页</h1>")
	})

	main.Post("/register", Controllers.Register)
	main.Post("/login", Controllers.Login)
}
