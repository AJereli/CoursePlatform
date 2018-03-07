package main


import (

	"CoursePlatform/Base"
)

const (
	DBAddress = "localhost:3306"
)

var (
	DBMetaAuth = Base.DBInfo{
		Login: "root",
		Pass: "root",
		DBName: "CoursePlatformDB",
	}
)


