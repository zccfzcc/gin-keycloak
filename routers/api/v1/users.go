package v1

import (
	"gin-keycloak/keycloakadmin"
	"gin-keycloak/pkg/app"
	"gin-keycloak/pkg/e"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	if keycloakadmin.Admingo == nil || !keycloakadmin.Admingo.VerifyAdminToken() {
		adminGuy, _ := keycloakadmin.InitAdmin()
		keycloakadmin.Admingo = adminGuy
	}

	token := c.Request.Header.Get("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	appG.Response(http.StatusOK, e.SUCCESS, keycloakadmin.Admingo.ExtractUUIDfromToken("OWLY", token))
}
