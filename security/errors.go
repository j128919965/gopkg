package security

import "github.com/j128919965/gopkg/errors"

var ErrDidNotLogin = &errors.BizError{
	Code:  1,
	Msg:   "没有检测到登录信息",
	Cause: nil,
}

var ErrWrongToken = &errors.BizError{
	Code:  2,
	Msg:   "登录信息无效",
	Cause: nil,
}

var ErrNeedFresh = &errors.BizError{
	Code:  3,
	Msg:   "令牌需要刷新",
	Cause: nil,
}

var ErrNeedLogin = &errors.BizError{
	Code:  4,
	Msg:   "需要重新登录",
	Cause: nil,
}
