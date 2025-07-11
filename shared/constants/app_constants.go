package constants

type Exception struct {
	Code           int `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int    `json:"httpStatusCode"`
}

// IServiceOutput<T> equivalent
type ServiceOutput[T any] struct {
	Message        string     `json:"message,omitempty"`
	OutputData     T          `json:"outputData,omitempty"`
	Exception      *Exception `json:"exception,omitempty"`
	HttpStatusCode int        `json:"httpStatusCode"`
	RespStatusCode int        `json:"respStatusCode"`
}

// Final API response structure
type ApiResponse[T any] struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}


const(Access_Token string = "access_token"
Refresh_Token string = "refresh_token")