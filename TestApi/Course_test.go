package TestApi

import (
"testing"
)



func TestAddCourse(t *testing.T) {
	req := TestRequest{Method: "POST", JSON: `{"name" : "CourseTest", "in_skills" : ["skill1", "skiil2"], "out_skills" : ["oskil", "oskill2", "0skill3"] }`,
		Url: ApiPrefix + "/course/add"}
		req.MakeTest(t)

}


func TestGetCourse(t *testing.T) {
	req := TestRequest{Method: "GET", JSON: "",
		Url: ApiPrefix + "/course/getAll"}
	req.MakeTest(t)
}