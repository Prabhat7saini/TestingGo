package service

import "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/service"
type userLoginService interface{
	GetUser()string
}


type UserLoginService struct{
	access *service.ServiceAccess
}


func NewUserLoginService(access *service.ServiceAccess) *UserLoginService {
	return &UserLoginService{
		access: access,
	}
}
func GetUser()string{
	return "user"
}