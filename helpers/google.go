package helpers

import (
	"context"
	"fmt"
	"goshaka/configs"

	"github.com/goccy/go-json"

	"google.golang.org/api/idtoken"
)

type IdentityToolkitTokenInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
	Aud           string `json:"aud"`
	Iss           string `json:"iss"`
}

func VerifyIdToken(ctx context.Context, idToken string) (*IdentityToolkitTokenInfo, error) {
	var googleClientId string = configs.GetEnv("GOOGLE_CLIENT_ID") // from credentials in the Google dev console
	var info IdentityToolkitTokenInfo

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		return &info, fmt.Errorf("cannot assign new validator")
	}

	payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientId)
	if err != nil {
		return &info, fmt.Errorf("cannot validate the google token")
	}

	// Marshaling the map into JSON
	b, err := json.Marshal(payload.Claims)
	if err != nil {
		return &info, fmt.Errorf("cannot marshal the google token")
	}

	// Unmarshaling the JSON into the struct
	if err := json.Unmarshal(b, &info); err != nil {
		return &info, fmt.Errorf("cannot unmarshal the google token")
	}

	if info.Aud != configs.GetEnv("GOOGLE_CLIENT_ID") {
		return &info, fmt.Errorf("invalid client ID")
	}

	if info.Iss != "accounts.google.com" && info.Iss != "https://accounts.google.com" {
		return &info, fmt.Errorf("invalid iss")
	}

	return &info, nil
}
