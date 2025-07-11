package service

import (
	sharedRedis "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/redis"
	"go.uber.org/zap"
)

type ServiceAccess struct {
	Redis  sharedRedis.Client
	Logger *zap.Logger
}
