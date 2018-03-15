package TestApi

import (
	"testing"
)


func TestAddComment (t *testing.T){
	req := TestRequest{Method: "POST", JSON: `{
		  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjI0MTU0MzcsImlzcyI6IkNvdXJzZVBsYXRmb3JtIiwibmJmIjoxNTIxMTE5NDM3LCJyYW5nIjowLCJ1c2VySWQiOjF9.A-diKpt9Fng0WmL-djIYDr60zz74HFVXM8uJY1evPrA	",
		  "comment": "WHAT a pretty comment/nforme/nfor us",
		  "UserLectionTask_id": 5
		}
		`,
		Url: ApiPrefix + "/lection/solution/comment/add"}
	req.MakeTest(t)
}