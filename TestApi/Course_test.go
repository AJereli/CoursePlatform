package TestApi

import (
"testing"
)



func TestAddCourse(t *testing.T) {
	var addCourseReq = TestRequest{Method: "POST", JSON: `{"name" : "CourseTest", "in_skills" : ["skill1", "skiil2"], "out_skills" : ["oskil", "oskill2", "0skill3"] }`,
		Url: "http://localhost:16001/course/add"}

	res, err := addCourseReq.MakeRequest(t)

	if err != nil {
		t.Error(err)
	}
	PrintBody(res, t)
}


func TestGetCourse(t *testing.T) {
	var addCourseReq = TestRequest{Method: "GET", JSON: "",
		Url: "http://localhost:16001/course/getAll"}

	res, err := addCourseReq.MakeRequest(t)

	if err != nil {
		t.Error(err)
	}
	PrintBody(res, t)
}