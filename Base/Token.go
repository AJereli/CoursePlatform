package Base


import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)




var hmacSampleSecret = []byte("asdhguczx1412313214jifh")

func CreateToken (userId string) string{
	tokenTime := time.Now().Unix() + ExpiresTime
	claims := jwt.MapClaims{"userId" : userId, "nbf": time.Now().Unix(), "exp": tokenTime,"iss": AppName}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims.Valid()
	toketString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		panic (err)
	}
	return toketString

}

func ParseToken(tokenString string)  TokenInfo {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		tokenInfo := TokenInfo {ISS: claims["iss"].(string), UserName: claims["user_name"].(string), Rang: claims["rang"].(int), ExpTime: claims["exp"].(float64), Nbf:claims["nbf"].(float64)}
		return tokenInfo
	} else {
		Log.Error(err.Error())
		panic(err)
	}

}
