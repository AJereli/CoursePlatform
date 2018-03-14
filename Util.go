package main

import (
	"database/sql"
	"CoursePlatform/Base"
	"strings"
)

func (this * Course) Populate(rows *sql.Rows) {
	var inSkill, outSkill string

	err := rows.Scan(&this.Id, &this.Name, &inSkill, &outSkill)
	this.InSkills = strings.Split(inSkill, "\n")
	this.OutSkill = strings.Split(outSkill, "\n")
	Base.CheckErr(err)
}

func GetCoursesFromRows(rows * sql.Rows ) []Course{
	var courses []Course

	for rows.Next(){
		var course Course
		course.Populate(rows)
		courses = append(courses, course)
	}
	return courses
}


func (this * LectionTask) Populate(rows *sql.Rows) {
	err := rows.Scan(&this.Id, &this.Task, &this.Answer,
		&this.IsTest, &this.TestAns, &this.LectionId, &this.CourseId)
	Base.CheckErr(err)
}

func GetLectionTasksFromRows(rows * sql.Rows ) []LectionTask{
	var lts []LectionTask

	for rows.Next(){
		var task LectionTask
		task.Populate(rows)
		lts = append(lts, task)
	}
	return lts
}


func (this * CourseTask) Populate(rows *sql.Rows) {
	err := rows.Scan(&this.Id, &this.Task, &this.Answer)
	Base.CheckErr(err)
}

func GetCourseTasksFromRows(rows * sql.Rows ) []CourseTask{
	var ct []CourseTask

	for rows.Next(){
		var task CourseTask
		task.Populate(rows)
		ct = append(ct, task)
	}
	return ct
}

func (this * Lection) Populate(rows *sql.Rows) {
	var notes string
	//TODO imgs
	err := rows.Scan(&this.Id, &this.Name, &this.Title, &this.Information, &notes, &this.CourseId)
	this.Notes = strings.Split(notes, "\n")

	Base.CheckErr(err)
}

func (this *LectionTaskSolution) Populate (rows *sql.Rows){
	err := rows.Scan(&this.Id, &this.Mark, &this.UserId, &this.LectionId, &this.LectionTaskId, &this.CourseId, &this.Answer)
	Base.CheckErr(err)
}

func GetLectionFromRows(rows * sql.Rows ) []Lection{
	var lections []Lection

	for rows.Next(){
		var lection Lection
		lection.Populate(rows)
		lections = append(lections, lection)
	}
	return lections
}

