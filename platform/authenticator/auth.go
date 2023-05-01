package authenticator

import (
	"context"
	"errors"
	. "example.com/go-boot/platform/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// New instantiates the *Authenticator.
func New() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		AppConfig.Oidc.Issuer,
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		RedirectURL:  AppConfig.Oidc.RedirectUrl,
		ClientID:     AppConfig.Oidc.ClientId,
		ClientSecret: AppConfig.Oidc.ClientSecret,
		Endpoint:     provider.Endpoint(),
		Scopes:       AppConfig.Oidc.Scopes,
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
