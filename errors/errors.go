package errors

type BizError struct {
	Code  int
	Msg   string
	Cause error
}

func (e *BizError) Error() string {
	if e.Cause == nil {
		return e.Msg
	}
	return e.Msg + "; 原因：" + e.Cause.Error()
}

var notFound = &BizError{
	Code: 404,
	Msg:  "未找到该记录",
}

func New(message string, code int) error {
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

func NotFound() *BizError {
	return notFound
}

func IllIllegalArgument(error error) *BizError {
	if err, ok := error.(*BizError); ok {
		return err
	}
	return &BizError{
		Msg:  error.Error(),
		Code: 400,
	}
}
