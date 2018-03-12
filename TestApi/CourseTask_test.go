package TestApi

import (
	"testing"
)



func TestAddCourseTask(t *testing.T) {
	//randStr := GetRandString(10)
	var req = TestRequest{Method: "POST", JSON: `{"task" : "This example of of some problem for some student", 
						"answer" : "This example of some solution for some problem for some student", "course_id" : 1}`,
		Url: "http://localhost:16001/course/task/add"}

	res, err := req.MakeRequest(t)

	if err != nil {
		t.Error(err)
	}
	PrintBody(res, t)
}


func TestGetAllCourseTask(t *testing.T) {
	var req = TestRequest{Method: "GET", JSON: "{\"course_id\" : 1}",
		Url: "http://localhost:16001/course/task/getAll"}

	res, err := req.MakeRequest(t)

	if err != nil {
		t.Error(err)
	}
	PrintBody(res, t)
}