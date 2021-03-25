package keycloakadmin

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Nerzal/gocloak/v7"
	"github.com/pkg/errors"
)

var keycloakAdminUser string
var keycloakAdminPass string
var keycloakURL string

func _loadEnvVars() {
	keycloakAdminUser = os.Getenv("KEYCLOAK_USER")
	keycloakAdminPass = os.Getenv("KEYCLOAK_PASSWORD")
	keycloakURL = os.Getenv("KEYCLOAK_URL")

}

func _checkConnection() {
	for {
		res, err := http.Get(keycloakURL)
		if err == nil && res != nil {
			if res.StatusCode == 200 {
				log.Println("Keycloak instance up, ready and reachable !")
				break
			}
		}
		time.Sleep(2 * time.Second)
		log.Println("Waiting for keycloak instance...")
		log.Println(err)
	}
}

// InitAdmin InitAdmin
func InitAdmin() (*AdminGuy, error) {
	_loadEnvVars()
	_checkConnection()
	client := gocloak.NewClient(keycloakURL)
	if client == nil {
		log.Fatal()
		return nil, errors.Errorf("Couldn't connect to Keycloak instance")
	}

	adminGuy, err := NewAdmin(keycloakAdminUser, keycloakAdminPass, client)

	return adminGuy, err
}
