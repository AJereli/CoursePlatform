package Base

import (
	"encoding/json"
	"net/http"
)

const  (
	successCode = 200
)
// STATUS
var (
	SuccessApiStatus = ApiStatus{StatusCode:successCode, Message: "Success!"}
)

type ApiStatus struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
}

func (apiStatus *ApiStatus) Send (w http.ResponseWriter){
	Log.Info(apiStatus.Message)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(apiStatus); err != nil {
		Log.Error(err)
		panic(err)
	}
}