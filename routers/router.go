package routers

import (
	"gin-keycloak/middleware/jwt"
	"gin-keycloak/routers/api"
	"gin-keycloak/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(CORSEnable())
	r.Use(gin.Recovery())
	r.GET("/", api.GeOk)
	apiv := r.Group("/api")
	apiv1 := r.Group("/api/v1")
	apiv.Use()
	{
		apiv.POST("/auth/login", api.GetAuth)
		apiv.POST("/auth/refresh", api.GetAuthRefreshToken)
	}
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/auth/logout", api.Logout)
		apiv1.GET("/user/me", v1.GetUser)
	}
	return r
}

func CORSEnable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.Status(200)
			return
		}
		c.Next()
	}
}
