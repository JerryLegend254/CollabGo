package auth

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthAuthenticator struct {
	config *oauth2.Config
	state  string
}

func NewOAuthAuthenticator(config *oauth2.Config, state string) *OAuthAuthenticator {
	return &OAuthAuthenticator{
		config: config,
		state:  state, // Use a dynamic state in production
	}
}

func GetSpotifyUserInfo(client *http.Client) (string, error) {
	resp, err := client.Get("https://api.spotify.com/v1/me")
	if err != nil {
		return "", fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error response from Spotify: %s", resp.Status)
	}

	var user struct {
		Name string `json:"display_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", fmt.Errorf("failed to decode user info: %w", err)
	}

	return user.Name, nil
}
func GenerateRandomState() string {
	return fmt.Sprintf("%d", rand.Intn(100000))
}
