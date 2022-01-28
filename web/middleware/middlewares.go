package middleware

import (
	"context"
	"github.com/j128919965/gopkg/security"
	"github.com/j128919965/gopkg/web/advice"
	resp "github.com/j128919965/gopkg/web/resp"
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

func AuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		tag := request.Context().Value("queTag")
		if tag!=nil {
			payload, err := security.NewPayLoadFromJsonContext(request.Context())
			if err != nil {
				advice.HandleError(w,err)
			}
			if err := payload.Valid() ; err != nil {
				advice.HandleError(w,err)
				return
			}
			request = request.WithContext(context.WithValue(request.Context(),"payload",payload))
		}
		next(w,request)
	}
}