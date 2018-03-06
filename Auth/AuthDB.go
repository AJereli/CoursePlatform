package Auth

import "fmt"

const (
	DBAddress = "localhost:3306"
)

var (
	DBMetaAuth = DBInfo{
		Login: "root",
		Pass: "root",
		DBName: "CoursePlatformDB",
	}
)


func (dbInfo  DBInfo) GetDataSourceName () string{
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbInfo.Login, dbInfo.Pass,DBAddress,dbInfo.DBName)
}
