package helpers

import (
	"github.com/kataras/iris/v12"
)

func GetStatusCode(err error) int {

	if err == nil {
		return iris.StatusOK
	}

	switch err {
	case ErrInternalServerError:
		return iris.StatusInternalServerError
	case ErrNotFound:
		return iris.StatusNotFound
	case ErrConflict:
		return iris.StatusConflict
	case ErrInvalidCredentials:
		return iris.StatusUnauthorized
	case ErrUserIsNotActive:
		return iris.StatusUnauthorized
	default:
		return iris.StatusInternalServerError
	}
}
