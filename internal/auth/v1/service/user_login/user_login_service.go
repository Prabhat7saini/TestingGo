package userloginservice

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/models"
	userloginrepository "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/repository/user_login"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/dto"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/service"
	// service "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/service"
)

// ---------- Service contract ----------

type UserLoginService interface {
	GetUser(ctx context.Context, payload *dto.LoginDto) (models.UserLogin, error)
}

// ---------- Concrete implementation ----------

type userLoginService struct {
	repo   userloginrepository.UserLoginRepositoryMethods
	access service.ServiceAccess
}

// New is a constructor that wires the repository into the service.
func NewUserLoginService(repo userloginrepository.UserLoginRepositoryMethods) UserLoginService {
	return &userLoginService{repo: repo}
}

// GetUser validates the input and delegates to the repository.
// Replace the repo call with your real implementation.
func (s *userLoginService) GetUser(
	ctx context.Context,
	payload *dto.LoginDto,
) (models.UserLogin, error) {

	if payload == nil {
		return models.UserLogin{}, fmt.Errorf("payload must not be nil")
	}

	// Example repository call (adjust to your repo’s signature):
	return models.UserLogin{}, errors.New("login service not implemented yet")
}
