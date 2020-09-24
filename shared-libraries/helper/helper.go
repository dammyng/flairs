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
	j, err := json.Marshall(data)
	if err != nil{
		DisplayAppError(w, err, "Json marshal error", http.StatusForbidden)
	}
	w.Write(j)
}

func DisplayApiError(w http.ResponseWriter, message, response_code string, data interface{}, code int) {
	errObj := struct {
		Message      string
		Data         interface{}
		ResponseCode string
	}{
		message,
		data,
		response_code,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errObj); err == nil {
		w.Write(j)
	}
}
