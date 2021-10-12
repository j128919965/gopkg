package middleware

import (
	"github.com/j128919965/gopkg/web/advice"
	resp "github.com/j128919965/gopkg/web/dto"
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