package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func LoadKeys(jwksURL string) (jwk.Set, error) {
	ctx := context.Background()
	ks, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		return nil, err
	}
	return ks, nil
}

func UserFromJWT(r *http.Request, authKeys jwk.Set) (string, error) {
	token, err := jwt.ParseRequest(r, jwt.WithKeySet(authKeys))
	if err != nil {
		return "", fmt.Errorf("parse request: %w", err)
	}
	userID, exists := token.Subject()
	if !exists {
		return "", fmt.Errorf("missing user id")
	}
	return userID, nil
}
