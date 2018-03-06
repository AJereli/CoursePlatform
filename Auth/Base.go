package Auth


import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
	"net/http"
	_"github.com/gorilla/mux"
)

const(
	LimitJSONRead = 1048576
)



func SendJson (w http.ResponseWriter, v interface{}){
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


func WraperLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Info(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
