package exception

import (
	"net/http"

	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants"
)

type ErrorCode string

const (
	USER_NOT_FOUND        = "USER_NOT_FOUND"
	USER_ALREADY_EXISTS   = "USER_ALREADY_EXISTS"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	INVALID_CREDENTIALS   = "INVALID_CREDENTIALS"
)

var ErrorCodeErrorMessage = map[ErrorCode]constants.Exception{
	USER_NOT_FOUND: {
		Code:           http.StatusNotFound,
		Message:        "User Not Found",
		HttpStatusCode: 404,
	},
	USER_ALREADY_EXISTS: {
		Code:           409,
		Message:        "User Already Exists",
		HttpStatusCode: 409,
	},
	INTERNAL_SERVER_ERROR: {
		Code:           500,
		Message:        "Internal Server Error",
		HttpStatusCode: 500,
	},
	INVALID_CREDENTIALS: {
		Code:           401,
		Message:        "Invalid Credentials",
		HttpStatusCode: 401,
	},
}

func GetException(code ErrorCode) *constants.Exception {
	if ex, ok := ErrorCodeErrorMessage[code]; ok {
		return &ex
	}
	return &constants.Exception{
		Code:           500,
		Message:        "Unknown error code",
		HttpStatusCode: 500,
	}
}
