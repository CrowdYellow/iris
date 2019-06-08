package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"iris/app/Controllers"
	"iris/app/Middleware"
)

func Api(app *iris.Application) {
	var main = corsSetting(app, "/api/")
	main.Post("/register", Controllers.Register)
	main.Post("/login", Controllers.Login)
	main.PartyFunc("/users", func(user router.Party) {
		user.Use(Middleware.JwtHandler().Serve)
		user.Get("/", Controllers.Me)
		user.Put("/name", Controllers.UpdateUsersNameById)
		user.Put("/password", Controllers.UpdateUserPasswordById)
	})
}
