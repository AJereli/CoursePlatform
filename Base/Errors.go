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

	someFieldIsEmptyCode = 140

	courseNameExistCode = 150
	courseNotFoundCode = 151

	tokenTimeOutCode = 160



	taskAlreadyCreatedCode = 170
	taskNotExistsCode = 171
	userTasNotExistsCode = 172

	lectionAlreadyCreatedCode = 180
	lectionNotExistsCode = 181


	answerExists = 190

	unprocessableEntityCode = 422
)



//ERRORS
var (
	WrongParamsApiErr            = ApiError{wrongRegParmsCode, wrongParams}
	NotExistUserNameApiErr       = ApiError{userNameNotExistsCode, "Sorry, this user name is not available"}
	LoginApiErr                  = ApiError{loginErrCode, "Wrong login or password"}

	TokenTimeOutApiErr	         = ApiError{tokenTimeOutCode, "Access token time out"}

	UnprocessableEntityApiErr    = ApiError{unprocessableEntityCode, "Unprocessablee entity"}

	SomeFieldIsEmptyApiErr 		 = ApiError{someFieldIsEmptyCode, "Some field is empty!"}

	CourseNameNotExistApiErr     = ApiError{courseNameExistCode, "Course with this name already created"}
	CourseNotFoundApiErr         = ApiError{courseNotFoundCode, "Course with this ID not found!"}

	TaskAlreadyCreatedApiErr     = ApiError{taskAlreadyCreatedCode, "This task already created!"}
	TaskNotExistsApiErr			 = ApiError{taskNotExistsCode, "Task with this id is not exists"}
	UserTaskNotExists			 = ApiError{userTasNotExistsCode, "Solution with this id is not exists"}


	TaskAnswerExistsApiErr       = ApiError{answerExists, "You have already given such a solution to the problem"}

	LectionAlreadyCreatedApiErr  = ApiError{lectionAlreadyCreatedCode, "Lection with this name already created!"}
	LectionNotExistsApiErr       = ApiError{lectionNotExistsCode, "Lection with this id not exists"}
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