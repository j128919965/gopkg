package security

import "github.com/j128919965/gopkg/errors"

var ErrDidNotLogin = &errors.BizError{
	Code:  1,
	Msg:   "auth:没有检测到登录信息",
	Cause: nil,
}

var ErrWrongToken = &errors.BizError{
	Code:  2,
	Msg:   "auth:登录信息无效",
	Cause: nil,
}

var ErrNeedFresh = &errors.BizError{
	Code:  3,
	Msg:   "auth:令牌需要刷新",
	Cause: nil,
}

var ErrNeedLogin = &errors.BizError{
	Code:  4,
	Msg:   "auth:需要重新登录",
	Cause: nil,
}
