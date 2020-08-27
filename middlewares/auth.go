package middlewares

import (
	"fmt"

	"github.com/astaxie/beego"
	beegoContext "github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
)

// CheckAuthentication middleware to check if incoming request is authenticated
func CheckAuthentication(beegoCtx *beegoContext.Context) {
	if beegoCtx.Request.Method != "OPTIONS" {
		authorizationHeader := beegoCtx.Request.Header.Get("Authorization")
		if len(authorizationHeader) < 7 {
			beegoCtx.Output.SetStatus(401)
			beegoCtx.Output.JSON(map[string]string{"message": "Invalid Token"}, true, true)
			return
		}
		bearerToken := authorizationHeader[7:]
		_, err := jwt.Parse(bearerToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(beego.AppConfig.String("api_secret")), nil
		})
		if err != nil {
			beegoCtx.Output.SetStatus(401)
			beegoCtx.Output.JSON(map[string]string{"message": "Invalid Token"}, true, true)
			return
		}
		// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// }
	}
}
