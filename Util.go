package main

import (
	"database/sql"
	"CoursePlatform/Base"
	"strings"
)

func (course * Course) PopulateCourseFromRow(rows *sql.Rows) {
	var inSkill, outSkill string

	err := rows.Scan(&course.Id, &course.Name, &inSkill, &outSkill)
	course.InSkills = strings.Split(inSkill, "\n")
	course.OutSkill = strings.Split(outSkill, "\n")
	Base.CheckErr(err)
}

func GetCoursesFromRows(rows * sql.Rows ) []Course{
	var courses []Course

	for rows.Next(){
		var course Course
		course.PopulateCourseFromRow(rows)
		courses = append(courses, course)
	}
	return courses
}


func (lectionTask * LectionTask) PopulateLectionTaskFromRow(rows *sql.Rows) {
	err := rows.Scan(&lectionTask.Id, &lectionTask.Task, &lectionTask.Answer,
		&lectionTask.IsTest, &lectionTask.TestAns, &lectionTask.LectionId, &lectionTask.CourseId)
	Base.CheckErr(err)
}

func GetLectionTasksFromRows(rows * sql.Rows ) []LectionTask{
	var lts []LectionTask

	for rows.Next(){
		var task LectionTask
		task.PopulateLectionTaskFromRow(rows)
		lts = append(lts, task)
	}
	return lts
}


func (courseTask * CourseTask) PopulateCourseTaskFromRow(rows *sql.Rows) {
	err := rows.Scan(&courseTask.Id, &courseTask.Task, &courseTask.Answer)
	Base.CheckErr(err)
}

func GetCourseTasksFromRows(rows * sql.Rows ) []CourseTask{
	var ct []CourseTask

	for rows.Next(){
		var task CourseTask
		task.PopulateCourseTaskFromRow(rows)
		ct = append(ct, task)
	}
	return ct
}

func (lection * Lection) PopulateLectionFromRow(rows *sql.Rows) {
	var notes string
	//TODO imgs
	err := rows.Scan(&lection.Id, &lection.Name, &lection.Title, &lection.Information, &notes, &lection.CourseId)
	lection.Notes = strings.Split(notes, "\n")

	Base.CheckErr(err)
}

func GetLectionFromRows(rows * sql.Rows ) []Lection{
	var lections []Lection

	for rows.Next(){
		var lection Lection
		lection.PopulateLectionFromRow(rows)
		lections = append(lections, lection)
	}
	return lections
}

