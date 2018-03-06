package Auth


type User struct {
	UserName, Password string

}

type DBInfo struct {
	Login, Pass, DBName string
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