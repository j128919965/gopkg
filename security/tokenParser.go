package security

import (
	"github.com/golang-jwt/jwt"
)

type TokenParser struct {
	acKey []byte
	rfKey []byte
}

func NewTokenParser(acKey, rfKey string) *TokenParser {
	return &TokenParser{
		acKey: []byte(acKey),
		rfKey: []byte(rfKey),
	}
}

func (parser *TokenParser) Parse(token string , key []byte) (*PayLoad , error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key,nil
	})

	if err != nil {
		return nil , ErrWrongToken
	}

	err = parse.Claims.Valid()
	if err != nil {
		return nil,err
	}

	pl,ok := parse.Claims.(*PayLoad)
	if !ok {
		return nil,ErrWrongToken
	}

	return pl,nil
}

func (p *TokenParser) ParseRefreshToken(refreshToken string) (*PayLoad , error) {
	return p.Parse(refreshToken,p.rfKey)
}