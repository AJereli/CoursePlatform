package main


type Course struct {
	Id int `json:"id"`
	Name string `json:"name"`
	InSkills []string `json:"in_skills"`
	OutSkill []string `json:"out_skills"`
}

type CoueseTask struct {
	Task string `json:"task"`
	Answer string `json:"answer"`
	CourseId int `json:"course_id"`
}

type Lection struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Information string `json:"information"`

	//TODO NO IMAGES

	Notes []string `json:"notes"`
	CourseId int `json:"course_id"`
}