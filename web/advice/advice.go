package advice

import (
	"encoding/json"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/web/resp"
	"net/http"
)

func HandleResult(w http.ResponseWriter, result interface{}, err error) {
	if err!=nil {
		if err,ok:=err.(*errors.BizError);ok {
			WriteJson(w,err.Code,resp.BizFailure(err))
			return
		}
		WriteJson(w,500,resp.ErrFailure(err))
		return
	}
	WriteJson(w,200,resp.Success(result))
}

func HandleError(w http.ResponseWriter, err error) {
	HandleResult(w,nil,err)
}

func WriteJson(w http.ResponseWriter,code int,data interface{})  {
	w.WriteHeader(code)
	bytes, _ := json.Marshal(data)
	w.Write(bytes)
}