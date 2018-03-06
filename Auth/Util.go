package Auth

import (
	"time"
	"net/url"
	"reflect"
)

func (token TokenInfo) CheckExpTime ()bool {
	if time.Now().Unix() >= int64(token.ExpTime) {
		return false
	}
	return true
}

func checkAuthParams (params url.Values) bool{
	isOk := true
	if len(params) != 2 {
		return false

	}

	keys := reflect.ValueOf(params).MapKeys()

	k1, k2 := keys[0].String(), keys[1].String()

	if k1 != "user_name" || k2 != "password"{
		isOk = false
	}

	if !isOk{
		//TODO some bug, false when params is normal
		log.Debug("what???")
	}

	return isOk
}
