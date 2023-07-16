package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/zakariawahyu/go-iris-news/modules/content"
	"github.com/zakariawahyu/go-iris-news/pkg/response"
	"net/http"
)

type ContentController struct {
	contentServices content.ContentServices
}

func NewContentController(serv content.ContentServices) ContentController {
	return ContentController{
		contentServices: serv,
	}
}

func (ctrl *ContentController) NewsRowAll(c iris.Context) {
	ctx := c.Request().Context()
	types := c.Params().Get("type")
	key := c.Params().Get("key")
	limit := c.URLParamIntDefault("limit", 30)
	offset := c.URLParamIntDefault("offset", 0)

	content := ctrl.contentServices.GetContentAllRow(ctx, types, key, limit, offset)

	c.JSON(&response.SuccessResponse{
		Success: true,
		Code:    http.StatusOK,
		Data:    content,
	})
}
