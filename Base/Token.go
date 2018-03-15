package Base


import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)




var hmacSampleSecret = []byte("asdhguczx1412313214jifh")

func CreateToken (user User) string{
	tokenTime := time.Now().Unix() + ExpiresTime
	claims := jwt.MapClaims{"userId" : user.UserId, "rang" : user.Rang, "nbf": time.Now().Unix(), "exp": tokenTime,"iss": AppName}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims.Valid()
	toketString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		panic (err)
	}
	return toketString

}

func (ti *TokenInfo) Populate(tokenString string)   {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})
	//if err != nil {
	//	Log.Error(err.Error())
	//	panic(err)
	//}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		ti.ISS     =   claims["iss"].(string)
		ti.UserId  =   claims["userId"].(float64)
		ti.Rang    =   int(claims["rang"].(float64))
		ti.ExpTime =   claims["exp"].(float64)
		ti.Nbf     =   claims["nbf"].(float64)}


		}


