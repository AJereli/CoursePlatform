package Base

import "fmt"

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`

}

type DBInfo struct {
	Login, Pass, DBName string
}

func (dbInfo DBInfo) GetDataSourceName (DBAddress string) string{
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbInfo.Login, dbInfo.Pass, DBAddress,dbInfo.DBName)
}


type TokenInfo struct {
	UserName string `json:"user_name"`
	Rang int `json:"rang"`
	ISS string `json:"iss"`
	ExpTime float64 `json:"exp"`
	Nbf float64 `json:"nbf"`

}

type JSONToken struct{
	AccessToken string `json:"access_token"`
}



