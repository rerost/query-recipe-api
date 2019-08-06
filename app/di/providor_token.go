package di

import "golang.org/x/oauth2"

func ProvidorToken(cfg Config) *oauth2.Token {
	return &oauth2.Token{AccessToken: cfg.AccessToken()}
}
