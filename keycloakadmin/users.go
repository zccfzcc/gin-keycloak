package keycloakadmin

import (
	"context"

	"github.com/Nerzal/gocloak/v7"
)


// AdminGuy AdminGuy
type AdminGuy struct {
	AccessToken  string
	RefreshToken string
	Client       gocloak.GoCloak
}

// NewAdmin NewAdmin
func NewAdmin(username string, password string, client gocloak.GoCloak) (*AdminGuy, error) {

	token, err := client.LoginAdmin(context.Background(), username, password, "master")
	if err != nil {
		return nil, err
	}

	return &AdminGuy{AccessToken: token.AccessToken, RefreshToken: token.RefreshToken, Client: client}, nil
}

// CreateUser CreateUser
func (adminGuy *AdminGuy) CreateUser(realm string, user gocloak.User) (string, error) {

	ID, err := adminGuy.Client.CreateUser(context.Background(), adminGuy.AccessToken, realm, user)
	return ID, err

}

// SetUserPassword SetUserPassword
func (adminGuy *AdminGuy) SetUserPassword(realm string, userID string, password string, tempo bool) error {
	return adminGuy.Client.SetPassword(context.Background(), adminGuy.AccessToken, userID, realm, password, tempo)
}

// SearchUserByUsername SearchUserByUsername
func (adminGuy *AdminGuy) SearchUserByUsername(realm string, username string) ([]*gocloak.User, error) {
	searchParams := gocloak.GetUsersParams{
		Search: &username,
	}
	return adminGuy.Client.GetUsers(context.Background(), adminGuy.AccessToken, realm, searchParams)
}

// GetUserByUUID GetUserByUUID
func (adminGuy *AdminGuy) GetUserByUUID(realm string, uuid string) (*gocloak.User, error) {

	return adminGuy.Client.GetUserByID(context.Background(), adminGuy.AccessToken, realm, uuid)

}
