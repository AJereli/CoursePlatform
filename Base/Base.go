package Base


import (
	"encoding/json"
	"io"
	"io/ioutil"

	"net/http"
	_"github.com/gorilla/mux"
)

const(
	LimitJSONRead = 1048576
)


const (
	AppName = "CoursePlatform"
	ExpiresTime = 60 * 60 * 24 * 15
)

func UnmarshalAndSend(w http.ResponseWriter, v interface{}){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func ReadRequestBody (r * http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitJSONRead))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	return body
}


func UnmarshalRequest (r * http.Request, v interface{}) error {
	body := ReadRequestBody(r)
	return json.Unmarshal(body, &v)
}


