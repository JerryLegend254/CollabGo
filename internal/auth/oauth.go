package auth

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"golang.org/x/oauth2"
)

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type OAuthAuthenticator struct {
	config *oauth2.Config
	state  string
}

func NewOAuthAuthenticator(config *oauth2.Config, state string) *OAuthAuthenticator {
	return &OAuthAuthenticator{
		config: config,
		state:  state, // TODO: Use a dynamic state in production
	}
}

func GetSpotifyUserInfo(client *http.Client) (*User, error) {
	resp, err := client.Get("https://api.spotify.com/v1/me")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response from Spotify: %s", resp.Status)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return &user, nil
}

func GenerateRandomState() string {
	return fmt.Sprintf("%d", rand.Intn(100000))
}
