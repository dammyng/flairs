package helper

import (
	"encoding/json"
	"net/http"
)

type(
	ErrorObj struct {
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus int `json:"status"`
	}

	HttpResponse struct{
		Message string `json:"message"`
		Code int `json:"code"`
			}
)

func WriteJsonResponse(w http.ResponseWriter, data interface{}, code int)  {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	j, err := json.Marshal(data)
	if err != nil{
		DisplayAppError(w, err, "Json marshal error", http.StatusForbidden)
	}
	w.Write(j)
}

func DisplayAppError(w http.ResponseWriter, err error, message string, code int) {
	errObj := ErrorObj{
		Error:      err.Error(),
		Message:    message,
		HttpStatus: code,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errObj); err == nil {
		w.Write(j)
	}
}
