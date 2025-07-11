package authhandler

import (


	"github.com/gin-gonic/gin"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/dto"
	authservice "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/service/auth_service"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants/exception"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/utils"
)

type AuthHandlerMethods interface {
	Login(ctx *gin.Context)
}

type authHandler struct {
	authService authservice.AuthServiceMethods
}

func NewAuthHandler(authService authservice.AuthServiceMethods) AuthHandlerMethods {
	return &authHandler{authService: authService}
}

func (ah *authHandler) Login(ctx *gin.Context) {
	var req dto.LoginDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// exc, _ := utils.CreateUserValidationErrors(err)


	
		resp := utils.ServiceError[struct{}](exception.USER_NOT_FOUND)
		utils.SendRestResponse(ctx, resp)
		return
	}

	output := ah.authService.Login(ctx, &req)
	utils.SendRestResponse(ctx, output)

}
