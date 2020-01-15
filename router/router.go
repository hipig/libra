package router

import (
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/controller"
)

func SetRouter(app *iris.Application) {
	app.AllowMethods(iris.MethodOptions)
	ctr := &controller.Controller{}

	api := app.Party("/api")
	api.Post("/", ctr.GenerNames)

}
