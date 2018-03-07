package main

import (
	"net/http"
	"CoursePlatform/Base"
	"io/ioutil"
	"io"
	"strings"
	"encoding/json"
	"database/sql"
	"github.com/prometheus/common/log"

)

func AddCourse (w http.ResponseWriter, r * http.Request){
	var course Course

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, Base.LimitJSONRead))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &course); err != nil {
		Base.UnprocessableEntityApiErr.Send(w)
		return
	}

	var inSkill, outSkill string

	inSkill = strings.Join(course.InSkills, "\n")
	outSkill = strings.Join(course.OutSkill, "\n")

	var courseNameNotExists bool
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)
	defer db.Close()

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM course WHERE name = ?)", course.Name).Scan(&courseNameNotExists)

	if courseNameNotExists {
		Base.CourseNameNotExistApiErr.Send(w)
		db.Close()
		return
	}


	stmt, err := db.Prepare("INSERT course SET name=?, InSkills=?, OutSkills=?")
	Base.CheckErr(err)


	res, err := stmt.Exec(course.Name, inSkill, outSkill)
	Base.CheckErr(err)
	log.Info(res)

	//if user.Password == truePassword{
	//	trueToken = Base.CreateToken(user.UserName)
	//	Base.SendJson(w, Base.JSONToken{AccessToken: trueToken})
	//	db.Query("UPDATE user SET access_token = ? WHERE name = ?", trueToken, user.UserName)
	//} else {
	//	Base.LoginApiErr.Send(w)
	//}
}

func GetCourses (w http.ResponseWriter, r * http.Request){
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	defer db.Close()
	Base.CheckErr(err)
	err = db.Ping()
	Base.CheckErr(err)

	rows, err := db.Query("SELECT id, name, InSkills, OutSkills FROM course")
	Base.CheckErr(err)


	msgs := GetCoursesFromRows(rows)
	defer rows.Close()

	Base.SendJson(w, msgs)
}

func AddCourseTask (w http.ResponseWriter, r * http.Request){

}

func GetCourseTasks (w http.ResponseWriter, r * http.Request){

}

func AddLection (w http.ResponseWriter, r * http.Request){

}

func GetCourseLections (w http.ResponseWriter, r * http.Request){

}

func AddLectionTask (w http.ResponseWriter, r * http.Request){

}

func GetLectionTasks (w http.ResponseWriter, r * http.Request){

}

