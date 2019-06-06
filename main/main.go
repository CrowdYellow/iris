package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"iris/config"
	"iris/routes"
)

func main()  {
	app := iris.New()
	app.Use(logger.New())
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		config.NotFound(ctx, "404 Not Found")
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		_, _ = ctx.WriteString("Oups something went wrong, try again")
	})
	routes.Api(app)
	_ = app.Run(iris.Addr(config.GO_LANG_PORT), iris.WithCharset("UTF-8"))
}