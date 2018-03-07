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