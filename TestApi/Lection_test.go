package TestApi

import "testing"

func TestAddLection (t *testing.T){
	req := TestRequest{Method: "POST", JSON: `{"name" : "Lection name", 
						"title" : "Lection title", "information" : "INFO INFO INFO\n info ninini\nINFO INFO INFO", "notes" : ["note 1", "note 2"], "course_id" : 1}`,
		Url: ApiPrefix + "/lection/add"}
	req.MakeTest(t)

}

func TestAddLectionTaskSolution (t * testing.T){
	req := TestRequest{Method: "POST", JSON: `{
		  "answer": "Sample ans",
		  "user_id": 1,
		  "lection_id": 1,
		  "course_id": 1,
		  "task_id": 1
		}
		`,
		Url: ApiPrefix + "/lection/solution/add"}
	req.MakeTest(t)
}
