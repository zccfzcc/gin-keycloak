package api

import (
	"gin-keycloak/pkg/app"
	"gin-keycloak/pkg/e"
	"gin-keycloak/service/authservice"

	"net/http"

	"github.com/astaxie/beego/validation"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	json := make(map[string]string) //注意该结构接受的内容
	c.BindJSON(&json)
	username := json["username"]
	password := json["password"]
	if username == "" {
		username, _ = c.GetPostForm("username")
		password, _ = c.GetPostForm("password")
	}

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	auth := authservice.Auth{
		Realm:    "OWLY",
		ClientID: "owly-cli",
		Username: username,
		Password: password,
	}
	appG.Response(http.StatusOK, e.SUCCESS, auth.Login())
}

func GetAuthRefreshToken(c *gin.Context) {
	appG := app.Gin{C: c}
	json := make(map[string]string) //注意该结构接受的内容
	c.BindJSON(&json)
	refreshToken := json["refresh_token"]
	if refreshToken == "" {
		refreshToken, _ = c.GetPostForm("refresh_token")
	}
	auth := authservice.Auth{
		Realm:    "OWLY",
		ClientID: "owly-cli",
		RefToken: refreshToken,
	}
	appG.Response(http.StatusOK, e.SUCCESS, auth.RefreshToken())
}
func Logout(c *gin.Context) {
	appG := app.Gin{C: c}
	json := make(map[string]string) //注意该结构接受的内容
	c.BindJSON(&json)
	refreshToken := json["refresh_token"]
	if refreshToken == "" {
		refreshToken, _ = c.GetPostForm("refresh_token")
	}

	auth := authservice.Auth{
		Realm:    "OWLY",
		ClientID: "owly-cli",
		RefToken: refreshToken,
	}
	appG.Response(http.StatusOK, e.SUCCESS, auth.Logout())
}

/**
  获取路由参数
*/
func GeOk(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, "OK")
}
