package authservice

import (
	"fmt"
	"gin-keycloak/keycloakadmin"
)

type Auth struct {
	Realm    string `json:"realm"`
	ClientID string `json:"clientID"`
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RefToken string `json:"refresh_token"`
}

func (a *Auth) Login() map[string]string {
	if keycloakadmin.Admingo == nil || !keycloakadmin.Admingo.VerifyAdminToken() {
		adminGuy, _ := keycloakadmin.InitAdmin()
		keycloakadmin.Admingo = adminGuy
	}
	jwt, _ := keycloakadmin.Admingo.Login(a.Realm, a.ClientID, a.Username, a.Password)
	token := map[string]string{}
	fmt.Println(jwt)
	if jwt != nil {
		token = map[string]string{
			"access_token":  jwt.AccessToken,
			"refresh_token": jwt.RefreshToken,
		}
	}

	return token
}

func (a *Auth) RefreshToken() map[string]string {
	if keycloakadmin.Admingo == nil || !keycloakadmin.Admingo.VerifyAdminToken() {
		adminGuy, _ := keycloakadmin.InitAdmin()
		keycloakadmin.Admingo = adminGuy
	}
	jwt, err := keycloakadmin.Admingo.GetRefreshToken(a.Realm, a.ClientID, a.RefToken)
	if err != nil {
		fmt.Println(err.Error())
	}
	token := map[string]string{}
	if jwt != nil {
		token = map[string]string{
			"access_token":  jwt.AccessToken,
			"refresh_token": jwt.RefreshToken,
		}
	}

	return token
}
func (a *Auth) Logout() string {
	if keycloakadmin.Admingo == nil || !keycloakadmin.Admingo.VerifyAdminToken() {
		adminGuy, _ := keycloakadmin.InitAdmin()
		keycloakadmin.Admingo = adminGuy
	}
	err := keycloakadmin.Admingo.Logout(a.Realm, a.ClientID, a.RefToken)
	if err != nil {
		return err.Error()

	}

	return "成功退出系统！"
}
