package keycloakadmin

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v7"
	"github.com/dgrijalva/jwt-go/v4"
)

// VerifyToken VerifyToken
func (adminGuy *AdminGuy) VerifyToken(realm string, clientID string, token string) (bool, error) {

	cliSecret, err := adminGuy.GetClientSecret(realm, clientID)
	if err != nil {
		return false, err
	}

	res, err := adminGuy.Client.RetrospectToken(context.Background(), token, clientID, cliSecret, realm)
	if err != nil {
		log.Printf("Fails at adminGuy.VerifyToken %v", err)
		return false, err
	}
	return *res.Active, nil
}

// VerifyToken VerifyToken
func (adminGuy *AdminGuy) VerifyAdminToken() bool {
	tokenActive, _ := adminGuy.VerifyToken("master", "master-realm", adminGuy.AccessToken)
	return tokenActive
}

// GetRefreshToken GetRefreshToken
func (adminGuy *AdminGuy) GetRefreshToken(realm string, clientID string, refreshToken string) (*gocloak.JWT, error) {
	cliSecret, err := adminGuy.GetClientSecret(realm, clientID)
	if err != nil {
		return nil, err
	}
	jwt, err := adminGuy.Client.RefreshToken(context.Background(), refreshToken, clientID, cliSecret, realm)
	if err != nil {
		log.Printf("Fails at adminGuy.RefreshToken %v", err)
		return nil, err
	}
	return jwt, nil
}

// Login Login
func (adminGuy *AdminGuy) Login(realm string, clientID string, username string, password string) (*gocloak.JWT, error) {

	cliSecret, err := adminGuy.GetClientSecret(realm, clientID)
	if err != nil {
		return nil, err
	}
	jwt, err := adminGuy.Client.Login(context.Background(), clientID, cliSecret, realm, username, password)
	if err != nil {
		log.Printf("Fails at adminGuy.Login %v", err)
		return nil, err
	}
	return jwt, nil
}
func (adminGuy *AdminGuy) Logout(realm string, clientID string, refreshToken string) error {
	cliSecret, err := adminGuy.GetClientSecret(realm, clientID)
	if err != nil {
		return err
	}
	return adminGuy.Client.Logout(context.Background(), clientID, cliSecret, realm, refreshToken)
}

// ExtractUUIDfromToken ExtractUUIDfromToken
func (adminGuy *AdminGuy) ExtractUUIDfromToken(realm string, token string) string {
	decoded, _, _ := adminGuy.Client.DecodeAccessToken(context.Background(), token, realm, "")
	claims := *decoded.Claims.(*jwt.MapClaims)
	return claims["name"].(string)
}
