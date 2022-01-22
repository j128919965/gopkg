package security

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
)

var defaultMaxCount int32 = 5

type PayLoad struct {
	QueTag    struct{} `json:"queTag"`
	UserId    int64    `json:"userId"`
	Expired   int64    `json:"expired"`
	Role      int32    `json:"role"`
	Refreshed int32    `json:"refreshed"`
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

func NewPayLoadFromJsonContext(ctx context.Context) (*PayLoad, error) {
	uid, err := strconv.ParseInt(ctx.Value("userId").(json.Number).String(), 10, 64)
	if err != nil {
		return nil, ErrWrongToken
	}
	expired, err := strconv.ParseInt(ctx.Value("expired").(json.Number).String(), 10, 64)
	if err != nil {
		return nil, ErrWrongToken
	}
	role, err := strconv.ParseInt(ctx.Value("role").(json.Number).String(), 10, 32)
	if err != nil {
		return nil, ErrWrongToken
	}
	refreshed, _ := strconv.ParseInt(ctx.Value("refreshed").(json.Number).String(), 10, 64)
	if err != nil {
		return nil, ErrWrongToken
	}
	return &PayLoad{
		QueTag:    struct{}{},
		UserId:    uid,
		Expired:   expired,
		Role:      int32(role),
		Refreshed: int32(refreshed),
	}, nil
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
