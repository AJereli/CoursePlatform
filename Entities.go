package main


type Course struct {
	Id int `json:"id"`
	Name string `json:"name"`
	InSkills []string `json:"in_skills"`
	OutSkill []string `json:"out_skills"`
}

type CourseTask struct {
	Id int `json:"id"`
	Task string `json:"task"`
	Answer string `json:"answer"`
	CourseId int `json:"course_id"`
}

type Lection struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
	Information string `json:"information"`

	//TODO: NO IMAGES

	Notes []string `json:"notes"`
	CourseId int `json:"course_id"`
}

func (self * Lection) IsReady () bool {
	return self.Name != "" && self.Title != "" && self.Information != "" && self.CourseId != 0
}

type LectionTask struct {
	Id int `json:"id"`
	Task string `json:"task"`
	Answer string `json:"answer"`
	IsTest bool `json:"isTest"`
	TestAns int  `json:"test_ans"`
	CourseId int `json:"course_id"`
	LectionId int `json:"lection_id"`
}

type LectionTaskSolution struct {
	Id int `json:"id"`
	Answer string `json:"answer"`
	UserId int `json:"user_id"`
	LectionId int  `json:"lection_id"`
	CourseId int `json:"course_id"`
	LectionTaskId int `json:"task_id"`
}