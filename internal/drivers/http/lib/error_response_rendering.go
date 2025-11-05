package lib

import (
	"net/http"
	"reflect"
	"retail_flow/internal/shared/exceptions"
	"strings"

	"github.com/go-chi/render"
)

func RenderErrorResponse(err error, w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var errStr = err.Error()

	if hasFieldInStruct(exceptions.BadRequestExceptions, errStr) {
		statusCode = 400
	} else if hasFieldInStruct(exceptions.UnauthenticatedExceptions, errStr) {
		statusCode = 401
	} else if hasFieldInStruct(exceptions.UnauthorizedExceptions, errStr) {
		statusCode = 403
	} else if hasFieldInStruct(exceptions.NotFoundExceptions, errStr) {
		statusCode = 404
	} else if hasFieldInStruct(exceptions.ConflictExceptions, errStr) {
		statusCode = 409
	} else {
		statusCode = 500
	}

	msg := err.Error()
	msg = strings.ToUpper(msg[0:1]) + msg[1:]

	render.Status(r, statusCode)
	render.JSON(w, r, map[string]string{
		"message": msg,
	})
}

func hasFieldInStruct(stc any, err string) bool {
	typeOf := reflect.TypeOf(stc)
	valueOf := reflect.ValueOf(stc)

	for i := 0; i < typeOf.NumField(); i++ {
		value := valueOf.Field(i).Interface().(error).Error()

		if strings.Contains(err, value) {
			return true
		}
	}

	return false
}
