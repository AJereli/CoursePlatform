package TestApi

import (
	"testing"
)



func TestAddCourseTask(t *testing.T) {
	req := TestRequest{Method: "POST", JSON: `{
			"task": "Sample task",
		  "answer": "Sample ans",
		  "isTest": false,
		  "test_ans": 0,
		  "course_id": 1,
		  "lection_id": 1
		}
		`,
		Url: ApiPrefix + "/lection/task/add"}
	req.MakeTest(t)
}


func TestAddLectionTask(t *testing.T) {
	req := TestRequest{Method: "POST", JSON: `{"task" : "This example of of some problem for some student", 
						"answer" : "This example of some solution for some problem for some student", "course_id" : 1}`,
		Url: ApiPrefix + "/course/task/add"}
	req.MakeTest(t)
}

func TestGetAllCourseTask(t *testing.T) {
	req :=TestRequest{Method: "GET", JSON: "{\"course_id\" : 1}",
		Url: ApiPrefix + "/course/task/getAll"}
	req.MakeTest(t)
}

func TestGetTaskByIdExists (t *testing.T){
	req := TestRequest{Method: "GET", JSON: "{\"lection_task_id\" : 1}",
		Url: ApiPrefix + "/lection/getById"}
	req.MakeTest(t)
}

func TestGetTaskByIdNotExists (t *testing.T){
	req := TestRequest{Method: "GET", JSON: "{\"lection_task_id\" : 100500}",
		Url: ApiPrefix + "/lection/getById"}
	req.MakeTest(t)
}

func TestGetTaskSolution (t *testing.T){
	req := TestRequest{Method: "GET", JSON: "{\"lection_task_solution_id\" : 1}",
		Url: ApiPrefix + "/lection/solution/getSolution"}
	req.MakeTest(t)
}

func TestEstimateTaskSolution (t *testing.T){
	json := `{"lection_task_solution_id" : 1, 
						"mark" : 5}`
	req := TestRequest{Method: "PUT", JSON : json, Url: ApiPrefix + "/lection/solution/estimate"}
	req.MakeTest(t)
}