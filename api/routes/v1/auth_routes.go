package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/intitalizer"
)

const (
	PublicApiV1    = "/public/v1/"
	ProtectedApiV1 = "/v1"
	PrivateApiV1   = "/private/v1"
)
const (
	AuthGroupPrefix  = "/auth"
	AuthProtectedApi = AuthGroupPrefix + ProtectedApiV1
	AuthPrivateApi   = AuthGroupPrefix + PrivateApiV1
	AuthPublicApi    = AuthGroupPrefix + PublicApiV1
)

type AuthRoutes struct {
}

func NewAuthRoutes(baseHandler *intitalizer.BaseHandler, routerGroup *gin.RouterGroup) *AuthRoutes {
	routerGroup.POST(AuthPublicApi+"/login", baseHandler.AuthHandler.Login)
	return &AuthRoutes{}
}
