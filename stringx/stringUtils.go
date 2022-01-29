package stringx

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

var mobileReg = regexp.MustCompile(`(?:0|86|\+86)?1[3-9]\d{9}`)

var emailReg = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)

var websiteReg = regexp.MustCompile(`(https://|http://)?([\w-]+\.)+[\w-]+(:\d+)*(/[\w- ./?%&=]*)?`)

func StartsWith(s,sub string) bool{
	l := len(sub)
	if l > len(s) {
		return false
	}

	for i := 0; i < l; i++ {
		if s[i] != sub[i] {
			return false
		}
	}

	return true
}

func IsBlank(str *string) bool {
	if str==nil {
		return true
	}
	if len(*str) == 0 {
		return true
	}
	if len(strings.TrimSpace(*str))==0 {
		return true
	}
	return false
}

func IsMobile(str *string) bool {
	return !IsBlank(str) && mobileReg.MatchString(*str)
}

func IsEmail(str *string) bool {
	return !IsBlank(str) && emailReg.MatchString(*str)
}

func IsGithub(str *string) bool {
	return !IsBlank(str) && StartsWith(*str,"https://github.com/")
}

func IsWebSite(str *string) bool {
	return !IsBlank(str) && websiteReg.MatchString(*str)
}

func GenerateNumCode(length int) string{
	arr := make([]byte,length)
	for i := 0; i < length; i++ {
		arr[i] = byte(rand.Intn(10)+'0')
		for i==0 && arr[i]=='0'{
			arr[i] = byte(rand.Intn(10)+'0')
		}
	}
	return string(arr)
}

func NullStringToPtr(v sql.NullString) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}

func Encrypt(salt, str string) string {
	dk, _ := scrypt.Key([]byte(str), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

func ValidPassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	var hasNumber, hasUpperCase, hasLowercase bool
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowercase = true
		}
	}
	return hasNumber && hasUpperCase && hasLowercase
}
