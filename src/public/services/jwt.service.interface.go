package services

import "github.com/sonuos1501/go-app/public/resources"

type IJwtService interface {
	Generate(payload *resources.UserResource) (string, error)
	Verify(tokenString string) (*UserClaims, error)
}
