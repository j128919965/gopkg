package stringx

import (
	"regexp"
	"strings"
)

var mobileReg = regexp.MustCompile(`(?:0|86|\\+86)?1[3-9]\\d{9}`)

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


