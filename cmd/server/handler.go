package server

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"github.com/zakariawahyu/go-iris-news/config"
	"github.com/zakariawahyu/go-iris-news/modules/content/controller"
	"github.com/zakariawahyu/go-iris-news/utils/exception"
	"log"
)

func NewAppHandler(i *iris.Application) {
	cfg := config.NewConfig()
	i.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"app":         cfg.App.Name,
			"version":     cfg.App.Version,
			"app_timeout": cfg.App.ContextTimeout,
		})
	})
}
func NewHandler(cfg *config.Config, serv *Services) {
	app := iris.New()

	app.Use(exception.ErrorHandler)

	NewAppHandler(app)
	v1 := app.Party("/v1")

	controller.NewContentController(v1, serv.contentServices)

	log.Fatal(app.Listen(viper.GetString("APP_ADDRESS")))
}
