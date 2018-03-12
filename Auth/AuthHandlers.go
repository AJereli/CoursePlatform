package Auth

import (
	"net/http"
	"database/sql"
	"io/ioutil"
	"io"
	_"github.com/go-sql-driver/mysql"
	"encoding/json"
	"CoursePlatform/Base"
	)





func Registration (w http.ResponseWriter, r * http.Request){
	//var regInfo Auth.RegistrationInfo

	params := r.URL.Query()


	if !checkAuthParams(params){
		Base.WrongParamsApiErr.Send(w)
		return
	}

	uID, userPass := params["name"][0], params["password"][0]

	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)

	err = db.Ping()
	Base.CheckErr(err)

	var userNameIsExists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM user WHERE name = ?)", uID).Scan(&userNameIsExists)

	if userNameIsExists {
		Base.NotExistUserNameApiErr.Send(w)
		db.Close()
		return
	}

	stmt, err := db.Prepare("INSERT user SET name=?, password=?")
	Base.CheckErr(err)

	res, err := stmt.Exec(uID, userPass)
	Base.CheckErr(err)

	log.Info(res)

	db.Close()

	Base.SuccessApiStatus.Send(w)
}

func Login (w http.ResponseWriter, r * http.Request){
	var user Base.User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, Base.LimitJSONRead))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}


	if err := json.Unmarshal(body, &user); err != nil {
		Base.UnprocessableEntityApiErr.Send(w)
	}

	var truePassword, trueToken string

	db, err := sql.Open("mysql", DBMetaAuth.GetDataSourceName(DBAddress))
	Base.CheckErr(err)

	db.QueryRow("SELECT password FROM user WHERE name = ?", user.UserName).Scan(&truePassword)
	defer db.Close()

	if user.Password == truePassword{
		trueToken = Base.CreateToken(user.UserName)
		Base.SendJson(w, Base.JSONToken{AccessToken: trueToken})
		db.Query("UPDATE user SET access_token = ? WHERE name = ?", trueToken, user.UserName)
	} else {
		Base.LoginApiErr.Send(w)
	}

}

