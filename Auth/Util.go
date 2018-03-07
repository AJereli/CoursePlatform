package Auth

import (
	"net/url"
	"reflect"

)

func checkAuthParams (params url.Values) bool{
	isOk := true
	if len(params) != 2 {
		return false

	}

	keys := reflect.ValueOf(params).MapKeys()

	k1, k2 := keys[0].String(), keys[1].String()

	if k1 != "name" || k2 != "password"{
		return false
	}

	if !isOk{
		//TODO some bug, false when params is normal
		log.Debug("what???")
	}

	return isOk
}
