package TestApi

import (
	"testing"
)


func TestUserRegistration (t *testing.T){
	req := TestRequest{Method: "GET", JSON: "",
		Url: "http://localhost:16000/registration?name=root&password=root"}
	req.MakeTest(t)
}

func TestUserLogin (t *testing.T) {
	req := TestRequest{Method: "GET", JSON: `{"user_name" : "root", "password" : "root" }`,
		Url: "http://localhost:16000/login"}
	req.MakeTest(t)
}