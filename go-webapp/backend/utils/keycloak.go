package utils

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v7"
	"log"
	"os"
)

type KeycloakClient struct {
	client gocloak.GoCloak
	realm  string
}

func NewKeycloakClient() *KeycloakClient {
	url := os.Getenv("KEYCLOAK_URL")
	realm := os.Getenv("KEYCLOAK_REALM")
	clientID := os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")

	client := gocloak.NewClient(url)

	// Authenticate with Keycloak
	token, err := client.LoginClient(context.Background(), clientID, clientSecret)
	if err != nil {
		log.Fatalf("Failed to login to Keycloak: %v", err)
	}

	return &KeycloakClient{
		client: *client,
		realm:  realm,
	}
}

func (kc *KeycloakClient) GetUserInfo(accessToken string) (*gocloak.UserInfo, error) {
	userInfo, err := kc.client.GetUserInfo(context.Background(), accessToken, kc.realm)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	return userInfo, nil
}

func (kc *KeycloakClient) ValidateToken(accessToken string) (bool, error) {
	_, err := kc.client.RetrospectiveToken(context.Background(), accessToken, kc.realm)
	if err != nil {
		return false, fmt.Errorf("invalid token: %w", err)
	}
	return true, nil
}