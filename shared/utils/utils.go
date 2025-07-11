package utils

import (
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants/exception"
)

func ServiceError[T any](code exception.ErrorCode) constants.ServiceOutput[T] {
	return constants.ServiceOutput[T]{
		Exception: exception.GetException(code),
	}
}
