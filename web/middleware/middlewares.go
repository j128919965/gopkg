package middleware

import (
	"gopkg/web/advice"
	resp "gopkg/web/dto"
	"net/http"
)

func ErrorHandler(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				advice.WriteJson(w,500,resp.MsgFailure(p.(string)))
			}
		}()
		next(w, r)
	}
}