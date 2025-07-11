package authservice

import (
	"context"

	"net/http"

	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/dto"
	authserviceaccess "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/service"
	userloginservice "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/service/user_login"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants/exception"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/utils"
	"gorm.io/gorm"
)

type AuthServiceMethods interface {
	Login(ctx context.Context, payload *dto.LoginDto) constants.ServiceOutput[string]
}

type authService struct {
	userLogin userloginservice.UserLoginServiceMethods
	access    *authserviceaccess.AuthServiceAccess
}

func NewAuthService(userLogin userloginservice.UserLoginServiceMethods, access *authserviceaccess.AuthServiceAccess) AuthServiceMethods {
	return &authService{
		userLogin: userLogin,
		access:    access,
	}
}

func (s *authService) Login(ctx context.Context, payload *dto.LoginDto) constants.ServiceOutput[string] {
	user, err := s.userLogin.GetUserByEmail(ctx, payload.Email, payload.PASSWORD)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
          s.access.Logger.Info("user not found")
			return utils.ServiceError[string](exception.USER_NOT_FOUND)
		}
		return utils.ServiceError[string](exception.INTERNAL_SERVER_ERROR)
	}
	// return "user.Token", nil

	// user := output.OutputData
     isvalid:=utils.CompareHashAndPassword(user.Password, payload.PASSWORD) 
	
	if !isvalid {
		return utils.ServiceError[string](exception.INVALID_CREDENTIALS)
	}

	return constants.ServiceOutput[string]{
		Message:        constants.LoginSuccess,
		OutputData:     "token",
		HttpStatusCode: http.StatusOK,
		RespStatusCode: http.StatusOK,
	}
}
