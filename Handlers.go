package main

import (
	"net/http"
	"CoursePlatform/Base"
	_ "io/ioutil"
	_ "io"
	"strings"
	"database/sql"
	"github.com/prometheus/common/log"
	"strconv"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	var course Course

	err := Base.UnmarshalRequest(r, &course)
	if err != nil {
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
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	defer db.Close()
	Base.CheckErr(err)
	err = db.Ping()
	Base.CheckErr(err)

	rows, err := db.Query("SELECT id, name, InSkills, OutSkills FROM course")
	Base.CheckErr(err)

	courses := GetCoursesFromRows(rows)
	defer rows.Close()

	Base.SendJson(w, courses)
}

func AddCourseTask(w http.ResponseWriter, r *http.Request) {
	var courseTask CourseTask

	err := Base.UnmarshalRequest(r, &courseTask)
	if err != nil {
		Base.UnprocessableEntityApiErr.Send(w)
		return
	}

	var courseIdExist, taskIsExist bool
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)
	defer db.Close()

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM course WHERE id = ?)", courseTask.CourseId).Scan(&courseIdExist)

	if !courseIdExist {
		Base.CourseNotFoundApiErr.Send(w)
		db.Close()
		return
	}

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM CourseTask WHERE task = ?)", courseTask.Task).Scan(&taskIsExist)

	if taskIsExist {
		Base.TaskAlreadyCreatedApiErr.Send(w)
		db.Close()
		return
	}

	stmt, err := db.Prepare("INSERT CourseTask SET task=?, answer=?, Course_id=?")
	Base.CheckErr(err)

	res, err := stmt.Exec(courseTask.Task, courseTask.Answer, courseTask.CourseId)
	Base.CheckErr(err)
	log.Info(res)

	var taskId int
	type Result struct {
		TaskId int `json:"task_id"`
	}
	db.QueryRow("SELECT id FROM CourseTask WHERE task = ?", courseTask.Task).Scan(&taskId)

	Base.SendJson(w, Result{TaskId: taskId})

}

func GetCourseTasks(w http.ResponseWriter, r *http.Request) {
	type CourseId struct {
		Id int `json:"course_id"`
	}

	var courseId CourseId
	err := Base.UnmarshalRequest(r, &courseId)
	if err != nil {
		Base.UnprocessableEntityApiErr.Send(w)
		return
	}

	var courseIdExist bool
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)
	defer db.Close()

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM course WHERE id = ?)", courseId.Id).Scan(&courseIdExist)

	if !courseIdExist {
		Base.CourseNotFoundApiErr.Send(w)
		db.Close()
		return
	}

	rows, err := db.Query("SELECT id, task, answer FROM CourseTask WHERE course_id=?", courseId.Id)
	Base.CheckErr(err)

	courseTasks := GetCourseTasksFromRows(rows)
	for i := 0; i < len(courseTasks); i++ {
		courseTasks[i].CourseId = courseId.Id
	}
	defer rows.Close()

	Base.SendJson(w, courseTasks)

}

func AddLection(w http.ResponseWriter, r *http.Request) {
	var lection Lection

	err := Base.UnmarshalRequest(r, &lection)
	if err != nil {
		Base.UnprocessableEntityApiErr.Send(w)
	}

	if !lection.IsReady() {
		Base.SomeFieldIsEmptyApiErr.Send(w)
		return
	}

	var lectionIsExist bool
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)
	defer db.Close()

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM lection WHERE name = ?)", lection.Name).Scan(&lectionIsExist)

	if lectionIsExist {
		Base.LectionAlreadyCreatedApiErr.Send(w)
		db.Close()
		return
	}

	//TODO images
	stmt, err := db.Prepare("INSERT Lection SET name=?, Title=?, information=?, notes=?, Course_id=?")
	Base.CheckErr(err)

	var notes string

	notes = strings.Join(lection.Notes, "\n")
	res, err := stmt.Exec(lection.Name, lection.Title, lection.Information, notes, lection.CourseId)
	Base.CheckErr(err)
	log.Info(res)

	var lectionId int
	type Result struct {
		TaskId int `json:"lection_id"`
	}
	db.QueryRow("SELECT id FROM Lection WHERE name = ?", lection.Name).Scan(&lectionId)

	Base.SendJson(w, Result{TaskId: lectionId})
}

func GetCourseLections(w http.ResponseWriter, r *http.Request) {
	type CourseId struct {
		Id int `json:"course_id"`
	}

	var courseId CourseId
	err := Base.UnmarshalRequest(r, &courseId)
	if err != nil {
		Base.UnprocessableEntityApiErr.Send(w)
		return
	}

	var courseIdExist bool
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)
	defer db.Close()

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM course WHERE id = ?)", courseId.Id).Scan(&courseIdExist)

	if !courseIdExist {
		Base.CourseNotFoundApiErr.Send(w)
		db.Close()
		return
	}

	rows, err := db.Query("SELECT id, name, title, information, notes, course_id FROM Lection WHERE course_id = ?", courseId.Id)
	Base.CheckErr(err)

	lections := GetLectionFromRows(rows)

	defer rows.Close()

	Base.SendJson(w, lections)

}

func AddLectionTask(w http.ResponseWriter, r *http.Request) {
	var lectionTask LectionTask

	err := Base.UnmarshalRequest(r, &lectionTask)
	if err != nil {
		Base.UnprocessableEntityApiErr.Send(w)
		return
	}

	var lectionIdExist, taskIsExist bool
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)
	defer db.Close()

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM lection WHERE id = ?)", lectionTask.LectionId).Scan(&lectionIdExist)

	if !lectionIdExist {
		Base.TaskAlreadyCreatedApiErr.Send(w)
		db.Close()
		return
	}

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM LectionTask WHERE task = ?)", lectionTask.Task).Scan(&taskIsExist)

	if taskIsExist {
		Base.TaskAlreadyCreatedApiErr.Send(w)
		db.Close()
		return
	}

	stmt, err := db.Prepare("INSERT LectionTask SET task=?, answer=?, isTest=?, testAns=?, Lection_id=?, Lection_Course_id=?")
	Base.CheckErr(err)

	res, err := stmt.Exec(lectionTask.Task, lectionTask.Answer, lectionTask.IsTest, lectionTask.TestAns, lectionTask.LectionId, lectionTask.CourseId)
	Base.CheckErr(err)
	log.Info(res)

	var taskId int
	type Result struct {
		TaskId int `json:"task_id"`
	}
	db.QueryRow("SELECT id FROM CourseTask WHERE task = ?", lectionTask.Task).Scan(&taskId)

	Base.SendJson(w, Result{TaskId: taskId})
}

func GetLectionTasks(w http.ResponseWriter, r *http.Request) {
	type LectionId struct {
		Id int `json:"lection_id"`
	}

	var lectionId LectionId
	Base.UnmarshalRequest(r, &lectionId)

	var courseIdExist bool
	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)
	defer db.Close()

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM lection WHERE id = ?)", lectionId.Id).Scan(&courseIdExist)

	if !courseIdExist {
		Base.LectionNotExistsApiErr.Send(w)
		db.Close()
		return
	}

	rows, err := db.Query("SELECT id, task, answer, isTest, testAns, Lection_id, Lection_Course_id FROM LectionTask WHERE Lection_id=?", lectionId.Id)
	Base.CheckErr(err)

	lectionTasks := GetLectionTasksFromRows(rows)

	defer rows.Close()

	Base.SendJson(w, lectionTasks)
}


func AddLectionTaskSolution(w http.ResponseWriter, r * http.Request){
	var solution LectionTaskSolution

	err := Base.UnmarshalRequest(r, &solution)
	if err != nil {
		Base.UnprocessableEntityApiErr.Send(w)
		return
	}

	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)
	defer db.Close()

	var answerExist, isTest bool
	var testAns, mark int

	db.QueryRow("SELECT EXISTS(SELECT 1 FROM UserLectionTask WHERE answer = ? AND user_id = ?)", solution.Answer, solution.UserId).Scan(&answerExist)

	if answerExist {
		Base.TaskAnswerExistsApiErr.Send(w)
		db.Close()
		return
	}

	db.QueryRow("SELECT isTest, testAns FROM LectionTask WHERE id =?", solution.LectionTaskId).Scan(&isTest, &testAns)

	var stmt *sql.Stmt
	if isTest{
		i, _ := strconv.Atoi(solution.Answer)
		if i == testAns{
			mark = 5
		}else {
			mark = 1
		}
	}

	stmt, err = db.Prepare("INSERT UserLectionTask SET mark=?, user_Id=?, answer=?, LectionTask_id=?, LectionTask_Lection_id=?, LectionTask_Lection_Course_id=?")
	Base.CheckErr(err)

	res, err := stmt.Exec(mark, solution.UserId, solution.Answer, solution.LectionTaskId, solution.LectionId, solution.CourseId)
	Base.CheckErr(err)
	log.Info(res)

	var solutionId int
	type Result struct {
		TaskId int `json:"solution_id"`
		Mark int `json:"mark"`
	}
	db.QueryRow("SELECT id FROM UserLectionTask WHERE answer = ? AND user_id = ?", solution.Answer, solution.UserId).Scan(&solutionId)

	Base.SendJson(w, Result{TaskId: solutionId, Mark: mark})
}

func GetLectionTaskSolutions(w http.ResponseWriter, r * http.Request){

}

func DeleteLectionTaskSolution (w http.ResponseWriter, r * http.Request){

}

func EstimateTaskSolution(w http.ResponseWriter, r * http.Request){

}