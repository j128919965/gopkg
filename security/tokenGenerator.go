package security

import (
	"github.com/golang-jwt/jwt"
	"github.com/j128919965/gopkg/errors"
	"time"
)

type TokenGenerator struct {
	acKey []byte
	rfKey []byte
}

func NewTokenGenerator(acKey, rfKey string) *TokenGenerator {
	return &TokenGenerator{
		acKey: []byte(acKey),
		rfKey: []byte(rfKey),
	}
}

func(generator *TokenGenerator) GenerateTokens(userId int64, role int32,refreshed int32) (*Tokens,error) {
	tokens := &Tokens{}
	signMethod := jwt.SigningMethodHS256
	now := time.Now()
	payLoad := &PayLoad{
		UserId:    userId,
		Expired:   now.Add(time.Hour * 24).Unix(),
		Role:      role,
		Refreshed: refreshed,
	}
	token := jwt.NewWithClaims(signMethod, payLoad)
	access, err := token.SignedString(generator.acKey)
	if err != nil {
		return nil,errors.Wrap(err,"生成AccessToken失败",0)
	}

	tokens.AccessToken = access

	payLoad.Expired = now.Add(time.Hour * 24 * 3).Unix()
	token = jwt.NewWithClaims(signMethod, payLoad)
	refresh, err := token.SignedString(generator.acKey)
	if err != nil {
		return nil,errors.Wrap(err,"生成RefreshToken失败",0)
	}

	tokens.RefreshToken = refresh
	return tokens,nil
}

func (generator *TokenGenerator) Refresh(payload *PayLoad) (*Tokens,error) {
	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return generator.GenerateTokens(payload.UserId, payload.Role, payload.Refreshed+1)
}
