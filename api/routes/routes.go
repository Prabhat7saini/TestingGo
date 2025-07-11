package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/cmd/
	v1 "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/api/routes/v1"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/intitalizer"
	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/cmd/app"
	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/cmd/app"
)

type Routes struct{}

func NewRoutes(router *gin.Engine, baseHandler *intitalizer.BaseHandler) {
  fmt.Println("inside roer")
	auths := router.Group("auth-service")
	 v1.NewAuthRoutes(baseHandler, auths)
	for _, r := range router.Routes() {
		fmt.Printf("Method: %s - Path: %s\n", r.Method, r.Path)
	}
	// app.addHealthPrivateRoutes(router, middlewares)
	// app.addHealthProtectedRoutes(router, middlewares)
	// app.addHealthPublicRoutes(router, middlewares)
	// app.addCatalogPrivateRoutes(router, middlewares)
	// app.addCatalogProtectedRoutes(router, middlewares)
	// app.addBillPaymentProtectedRoutes(router, middlewares)
	// app.addOrderProtectedRoutes(router, middlewares)
	// app.addCallBackPublicRoutes(router)
	// app.addRefundProtectedRoutes(router, middlewares)
	// return &Routes{}
}
