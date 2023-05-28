package exception

import (
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zakariawahyu/go-iris-news/helpers"
	"github.com/zakariawahyu/go-iris-news/utils/response"
)

func ErrorHandler(ctx iris.Context) {
	recorder := ctx.Recorder()

	defer func() {
		var err error
		code := iris.StatusInternalServerError

		if v := recover(); v != nil { // panic
			if panicErr, ok := v.(error); ok {
				err = panicErr
			} else {
				err = errors.New(fmt.Sprint(v))
			}
		} else {
			err = ctx.GetErr()
		}

		if err != nil {
			recorder.ResetBody()
			recorder.ResetHeaders()

			code = helpers.GetStatusCode(err)

			ctx.StopWithJSON(code, response.ErrorResponse{
				Success: false,
				Code:    code,
				Errors:  err.Error(),
			})
		}
	}()

	ctx.Next()
}
