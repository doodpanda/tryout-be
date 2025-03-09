package common

type ErrorResponse struct {
	OK    bool   `json:"ok" example:"false"`
	Error string `json:"error"`
}

type GeneralSuccessResponse struct {
	OK      bool   `json:"ok" example:"true"`
	Message string `json:"message"`
}

type GeneralSuccessHasDataResponse struct {
	OK      bool        `json:"ok" example:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		OK:    false,
		Error: err.Error(),
	}
}

func CreateGeneralSuccessResponse(msg string) *GeneralSuccessResponse {
	return &GeneralSuccessResponse{
		OK:      true,
		Message: msg,
	}
}

func CreateGeneralSuccessHasDataResponse(msg string, data interface{}) *GeneralSuccessHasDataResponse {
	return &GeneralSuccessHasDataResponse{
		OK:      true,
		Message: msg,
		Data:    data,
	}
}
