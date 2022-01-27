package advice

import (
	"encoding/json"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/stringx"
	"github.com/j128919965/gopkg/web/resp"
	"net/http"
	"strings"
)

func HandleResult(w http.ResponseWriter, result interface{}, err error) {
	if err!=nil {
		if err,ok:=err.(*errors.BizError);ok {
			WriteJson(w,400,resp.BizFailure(err))
			return
		}
		if stringx.StartsWith(err.Error(), "rpc error:") {
			str := err.Error()
			idx := strings.Index(str,"desc = ")
			if idx >= 0{
				idx += 7
				WriteJson(w,400,resp.MsgFailure(str[idx:]))
				return
			}
		}
		WriteJson(w,500,resp.ErrFailure(err))
		return
	}
	WriteJson(w,200,resp.Success(result))
}

func HandleError(w http.ResponseWriter, err error) {
	HandleResult(w,nil,err)
}

var contentType = "Content-Type"
var applicationJson = "application/json"

func WriteJson(w http.ResponseWriter,code int,data interface{})  {
	w.WriteHeader(code)
	w.Header().Set(contentType, applicationJson)
	bytes, _ := json.Marshal(data)
	w.Write(bytes)
}