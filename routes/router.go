package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func Routes(app *iris.Application)  {
	Web(app)
	Api(app)
}

func corsSetting(app *iris.Application, suffix string) (main iris.Party) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	/* 定义路由 */
	main = app.Party(suffix, crs).AllowMethods(iris.MethodOptions)
	return main
}