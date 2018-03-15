package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"CoursePlatform/Base"
)

type Route struct{
	Name, Method, Pattern string
	HandlerFunc http.HandlerFunc

}

type Routes []Route

type HandlFunc func(w http.ResponseWriter, r *http.Request)

var routes = Routes{
	Route{
		"AddCourse",
		"POST",
		"/course/add",
		AddCourse,
	},
	Route{
		"GetAllCourses",
		"GET",
		"/course/getAll",
		GetCourses,
	},
	Route{
		"GetAllCourseTasks",
		"GET",
		"/course/task/getAll",
		GetCourseTasks,
	},
	Route{
		"AddCourseTask",
		"POST",
		"/course/task/add",
		AddCourseTask,
	},
	Route{
		"AddLection",
		"POST",
		"/lection/add",
		AddLection,
	},
	Route{
		"GetCourseLections",
		"GET",
		"/lection/getAll",
		GetCourseLections,
	},
	Route{
		"GetAllLectionTasks",
		"GET",
		"/lection/task/getAll",
		GetLectionTasks,
	},
	Route{
		"AddLectionTask",
		"POST",
		"/lection/task/add",
		AddLectionTask,
	},
	Route{
		"GetLectionTask",
		"GET",
		"/lection/getById",
		GetLectionTaskById,
	},
	Route{
		"AddLectionTaskSolution",
		"POST",
		"/lection/solution/add",
		AddLectionTaskSolution,
	},

	Route{
		"GetLectionTaskSolution",
		"GET",
		"/lection/solution/getSolution",
		GetLectionTaskSolution,
	},
	Route{
		"DeleteLectionTaskSolution",
		"DELETE",
		"/lection/solution/delete",
		DeleteLectionTaskSolution,
	},
	Route{
		"EstimateTaskSolution",
		"PUT",
		"/lection/solution/estimate",
		EstimateTaskSolution,
	},
	Route{
		"AddComment",
		"POST",
		"/lection/solution/comment/add",
		AddComment,
	},
}

func InitRouter () *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes{
		var handler http.Handler
		handler = route.HandlerFunc
		handler = WraperLogger(handler, route.Name)

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}


func WraperLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		Base.Log.Info(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
