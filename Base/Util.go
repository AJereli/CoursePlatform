package Base

import (
	"time"

)

func (token TokenInfo) CheckExpTime ()bool {
	if time.Now().Unix() >= int64(token.ExpTime) {
		return false
	}
	return true
}

