package resp

import "github.com/j128919965/gopkg/errors"

type ApiResponse struct {
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Success bool `json:"success"`
}

func Success(data interface{}) *ApiResponse {
	return &ApiResponse{Data: data,Success: true}
}

func BizFailure(err *errors.BizError) *ApiResponse {
	return &ApiResponse{Data: err.Code,Message: err.Error() ,Success: false}
}

func ErrFailure(err error) *ApiResponse {
	return &ApiResponse{Success: false,Message: err.Error()}
}

func MsgFailure(msg string) *ApiResponse {
	return &ApiResponse{Success: false,Message: msg}
}
