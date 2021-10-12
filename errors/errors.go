package errors

type BizError struct {
	Code  int
	Msg   string
	Cause error
}

func (e *BizError) Error() string {
	return e.Msg
}

var NotFound = &BizError{
	Code: 404,
	Msg:  "未找到该记录",
}

func New(message string,code int) error {
	return &BizError{
		Msg:  message,
		Code: code,
	}
}

func Wrap(cause error, msg string, code int) error {
	return &BizError{
		Msg:   msg,
		Code:  code,
		Cause: cause,
	}
}
