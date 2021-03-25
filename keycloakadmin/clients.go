package keycloakadmin

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v7"
)

var Admingo *AdminGuy

// GetClientSecret GetClientSecret
func (adminGuy *AdminGuy) GetClientSecret(realm string, clientName string) (string, error) {

	var cli_id string

	res1, err := adminGuy.GetAllClients(realm)
	if err != nil {
		return "", err
	}

	for _, cli := range res1 {
		if clientID := cli.ClientID; *clientID == clientName {
			cli_id = *cli.ID
		}
	}

	res, err := adminGuy.Client.GetClientSecret(context.Background(), adminGuy.AccessToken, realm, cli_id)
	if err != nil {
		log.Printf("Error @ GetClientSecret: %v", err)
	}

	return *res.Value, err
}

// GetAllClients GetAllClients
func (adminGuy *AdminGuy) GetAllClients(realm string) ([]*gocloak.Client, error) {
	bill := true
	return adminGuy.Client.GetClients(context.Background(), adminGuy.AccessToken, realm, gocloak.GetClientsParams{ViewableOnly: &bill})
}

// GetClientId GetClientId
func (adminGuy *AdminGuy) GetClientId(realm string, clientName string) (string, error) {
	var cli_id string

	res1, err := adminGuy.GetAllClients(realm)
	if err != nil {
		return "", err
	}

	for _, cli := range res1 {
		if clientID := cli.ClientID; *clientID == clientName {
			cli_id = *cli.ID
		}
	}
	return cli_id, nil
}
