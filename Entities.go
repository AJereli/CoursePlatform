package main


type Course struct {
	Id 		   int      `json:"id"`
	Name       string   `json:"name"`
	InSkills   []string `json:"in_skills"`
	OutSkill   []string `json:"out_skills"`
}

type CourseTask struct {
	Id       int    `json:"id"`
	Task     string `json:"task"`
	Answer   string `json:"answer"`
	CourseId int    `json:"course_id"`
}

type Lection struct {
	Id   	    int		`json:"id"`
	Name  	    string  `json:"name"`
	Title 		string  `json:"title"`
	Information string  `json:"information"`

	//TODO: NO IMAGES

	Notes 		[]string `json:"notes"`
	CourseId 	int 	 `json:"course_id"`
}

func (this * Lection) IsReady () bool {
	return this.Name != "" && this.Title != "" && this.Information != "" && this.CourseId != 0
}

type LectionTask struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Answer    string `json:"answer"`
	IsTest    bool   `json:"isTest"`
	TestAns   int    `json:"test_ans"`
	CourseId  int    `json:"course_id"`
	LectionId int    `json:"lection_id"`
}

type LectionTaskSolution struct {
	Id            int    `json:"id"`
	Mark          int    `json:"mark"`
	Answer        string `json:"answer"`
	UserId        int    `json:"user_id"`
	LectionId     int    `json:"lection_id"`
	CourseId      int    `json:"course_id"`
	LectionTaskId int    `json:"task_id"`
}


type Comment struct {
	Id                int    `json:"id"`
	Comment           string `json:"comment"`
	Date              int64    `json:"date"`
	UserId            int    `json:"user_id"`
	UserLectionTaskId int    `json:"UserLectionTask_id"`
	LectionId         int    `json:"Lection_id"`
	CourseId          int    `json:"Course_id"`
	LectionTaskId     int    `json:"LectionTask_id"`
}