package Base

import (
	"encoding/json"
	"net/http"
)

/// API error codes
const (
	wrongRegParmsCode = 100
	wrongParams = "Wrong params"

	userNameNotExistsCode = 101

	loginErrCode = 102



	tokenTimeOutCode = 160


	unprocessableEntityCode = 422


	courseNameExistCode = 150
)

var (
	WrongParamsApiErr         = ApiError{wrongRegParmsCode, wrongParams}
	NotExistUserNameApiErr    = ApiError{userNameNotExistsCode, "Sorry, this user name is not available"}
	LoginApiErr               = ApiError{loginErrCode, "Wrong login or password"}

	TokenTimeOutApiErr	      = ApiError{tokenTimeOutCode, "Access token time out"}

	UnprocessableEntityApiErr = ApiError{unprocessableEntityCode, "Unprocessablee entity"}

	CourseNameNotExistApiErr  = ApiError{courseNameExistCode, "Course with this name already created"}
)

type ApiError struct {
	ErrorCode int `json:"error_code"`
	Message string `json:"message"`
}

func (apiErr *ApiError) Send (w http.ResponseWriter){
	Log.Info(apiErr.Message)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(apiErr); err != nil {
		Log.Error(err)
		panic(err)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}