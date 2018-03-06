package Auth

import (
	_"awesomeProject/MiddelewareChain"
	"net/http"
	"database/sql"
	"awesomeProject/Auth"
	"io/ioutil"
	"io"
	"encoding/json"
)


const (
	AppName = "CoursePlatform"
	ExpiresTime = 60 * 60 * 24 * 15
)


func Registration (w http.ResponseWriter, r * http.Request){
	//var regInfo Auth.RegistrationInfo

	params := r.URL.Query()


	if !checkAuthParams(params){
		wrongParamsApiErr.send(w)
		return
	}

	uID, userPass := params["user_name"][0], params["password"][0]

	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName())
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	var userNameIsExists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE userid = ?)", uID).Scan(&userNameIsExists)

	if userNameIsExists {
		notExistUserNameApiErr.send(w)
		db.Close()
		return
	}

	accessToken := Auth.CreateToken(uID)

	stmt, err := db.Prepare("INSERT users SET userid=?, password=?, access_token=?")
	checkErr(err)

	res, err := stmt.Exec(uID, userPass, accessToken)
	checkErr(err)

	log.Info(res)

	db.Close()

	jsonToken := Auth.JSONToken{AccessToken: accessToken}

	SendJson(w, jsonToken)
}

func Login (w http.ResponseWriter, r * http.Request){
	var user User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitJSONRead))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}


	if err := json.Unmarshal(body, &user); err != nil {
		unprocessableEntityApiErr.send(w)
	}

	var truePassword, trueToken string

	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName())
	checkErr(err)

	db.QueryRow("SELECT password FROM users WHERE userid = ?", user.UserName).Scan(&truePassword)
	defer db.Close()

	if user.Password == truePassword{
		trueToken = Auth.CreateToken(user.UserName)
		SendJson(w, Auth.JSONToken{AccessToken: trueToken})
		db.Query("UPDATE users SET access_token = ? WHERE userid = ?", trueToken, user.UserName)
	} else {
		loginApiErr.send(w)
	}

}
