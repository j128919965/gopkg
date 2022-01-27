package resp

import (
	"fmt"
	"github.com/j128919965/gopkg/errors"
	"golang.org/x/crypto/scrypt"
)

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

func Encrypt(salt, str string) string {
	dk, _ := scrypt.Key([]byte(str), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}