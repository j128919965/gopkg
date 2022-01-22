package security

import "time"

var defaultMaxCount int32 = 5

type PayLoad struct {
	UserId    int64 `json:"userId"`
	Expired   int64 `json:"expired"`
	Role      int32 `json:"role"`
	Refreshed int32 `json:"refreshed"`
}

func (p *PayLoad) Valid() error {
	now := time.Now().Unix()

	if p.Role < 0 {
		return ErrWrongToken
	}

	if now < p.Expired {
		if p.Refreshed < defaultMaxCount {
			return ErrNeedFresh
		}
		return ErrNeedLogin
	}

	return nil
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

