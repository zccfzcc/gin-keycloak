package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"gin-keycloak/keycloakadmin"
	"gin-keycloak/pkg/e"

	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		if keycloakadmin.Admingo == nil || !keycloakadmin.Admingo.VerifyAdminToken() {
			adminGuy, _ := keycloakadmin.InitAdmin()
			keycloakadmin.Admingo = adminGuy
		}

		code = e.SUCCESS

		token := c.Request.Header.Get("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")
		fmt.Println(token)
		if token == "" {
			code = e.INVALID_PARAMS
		} else {

			verify, _ := keycloakadmin.Admingo.VerifyToken("OWLY", "owly-cli", token)

			if !verify {

				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
