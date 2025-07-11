package service

import userLoginService "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/service/user_login"
type BaseService struct{
	userLogin userLoginService.UserLoginService
}