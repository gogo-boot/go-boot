package authenticator

import (
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	. "gogo-boot/go-boot/platform/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// New instantiates the *Authenticator.
func NewOidc() (*Authenticator, error) {
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

// New instantiates the *Authenticator.
func NewOAuth2() (*Authenticator, error) {

	conf := oauth2.Config{
		RedirectURL:  AppConfig.Oauth2.RedirectUrl,
		ClientID:     AppConfig.Oauth2.ClientId,
		ClientSecret: AppConfig.Oauth2.ClientSecret,
		Endpoint:     endpoints.AzureAD(AppConfig.Oauth2.Tenant),
		Scopes:       AppConfig.Oauth2.Scopes,
	}

	return &Authenticator{
		Config: conf,
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
