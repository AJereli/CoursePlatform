package TestApi

import (
	"io"
	"os"
	"net/http"
	"testing"
	"strings"
	"fmt"
	"time"
	"math/rand"
)

type TestRequest struct {
	Method string
	JSON string
	Url string
}

func (testReq * TestRequest) MakeRequest(t *testing.T) (*http.Response, error){
	var request *http.Request
	var err error
	if testReq.JSON != ""{
		reader := strings.NewReader(testReq.JSON)

		request, err = http.NewRequest(testReq.Method, testReq.Url, reader)
	}else {
		request, err = http.NewRequest(testReq.Method, testReq.Url, nil)
	}

	if err != nil{
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(request)

	return res, err
}

func PrintBody (res *http.Response, t *testing.T){
	fmt.Println("Out for ---- \n      ",res.Request.URL, "\n")
	defer res.Body.Close()
	_, err := io.Copy(os.Stdout, res.Body)
	fmt.Println("\n--------------------------")
	if err != nil {
		t.Error(err)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func InitRand() {
	rand.Seed(time.Now().UnixNano())
}


func GetRandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		InitRand()
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}