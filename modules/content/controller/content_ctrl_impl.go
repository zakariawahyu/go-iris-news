package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/zakariawahyu/go-iris-news/modules/content/services"
	"github.com/zakariawahyu/go-iris-news/utils/response"
	"net/http"
)

type ContentControllerImpl struct {
	contentServices services.ContentServices
}

func NewContentController(e iris.Party, serv services.ContentServices) {
	content := &ContentControllerImpl{
		contentServices: serv,
	}
	e.Get("/read/:slug", content.GetBySlug)
}

func (ctrl *ContentControllerImpl) GetBySlug(c iris.Context) {
	slug := c.Params().Get("slug")

	ctx := c.Request().Context()

	content := ctrl.contentServices.GetBySlug(ctx, slug)

	c.JSON(&response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Data:    content,
	})
}
