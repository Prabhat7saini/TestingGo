package utils

import (

	"github.com/gin-gonic/gin"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants"
)

func SendRestResponse[T any](ctx *gin.Context, output constants.ServiceOutput[T]) {
	if output.Exception != nil {
		ctx.JSON(output.Exception.HttpStatusCode, constants.ApiResponse[any]{
			Code:    output.Exception.Code,
			Message: fallbackIfEmpty(output.Message, output.Exception.Message),
		})
		return
	}
	// if output.HttpStatusCode == 0 {
	// 	panic("HttpStatusCode is required and cannot be 0")
	// }
	// if output.RespStatusCode == 0 {
	// 	panic("RespStatusCode is required and cannot be 0")
	// }

	ctx.JSON(output.HttpStatusCode, constants.ApiResponse[T]{
		Code:    output.RespStatusCode,
		Message: fallbackIfEmpty(output.Message, "SUCCESS"),
		Data:    output.OutputData,
	})
}

func fallbackIfEmpty(preferred string, fallback string) string {
	if preferred != "" {
		return preferred
	}
	return fallback
}
