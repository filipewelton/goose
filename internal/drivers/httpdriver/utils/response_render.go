package utils

import (
	e "errors"
	"fmt"
	"reflect"
	"retail_workflow/internal/shared/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func Render(err error, ctx *gin.Context) {
	errStr := err.Error()
	statusCode := getStatusCodeFromError(err)
	msg := fmt.Sprint(strings.ToUpper(errStr[:1]), errStr[1:])

	ctx.JSON(statusCode, gin.H{
		"message": msg,
	})
}

func getStatusCodeFromError(err error) int {
	switch {
	case checkIfTheStructHasTheValue(errors.Err400, err):
		return 400
	case checkIfTheStructHasTheValue(errors.Error401, err):
		return 401
	case checkIfTheStructHasTheValue(errors.Err404, err):
		return 404
	default:
		return 500
	}
}

func checkIfTheStructHasTheValue(str any, err error) bool {
	valueOf := reflect.ValueOf(str)

	for i := 0; i < valueOf.NumField(); i++ {
		fieldValue := valueOf.Field(i).Interface()

		if e.Is(err, fieldValue.(error)) {
			return true
		}
	}

	return false
}
