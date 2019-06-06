package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func main()  {
	app := iris.New()
	app.Use(logger.New())
}