package server

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"github.com/zakariawahyu/go-iris-news/config"
	"github.com/zakariawahyu/go-iris-news/modules/content/controller"
	"github.com/zakariawahyu/go-iris-news/pkg/exception"
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

	contentController := controller.NewContentController(serv.contentServices)

	v2 := app.Party("/v2")

	v2.Get("/news-row", contentController.NewsRowAll)
	v2.Get("/news-row/:type/:key", contentController.NewsRowAll)

	log.Fatal(app.Listen(viper.GetString("APP_ADDRESS")))
}
