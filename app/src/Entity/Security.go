package Entity

import (
	"time"
)

type AccessToken struct {
	Identifier     string
	ClientId       string
	Expiry         time.Time
	UserIdentifier string
}

type Client struct {
	Identifier   string
	Name         string
	Secret       string
	RedirectURLs string
	Grants       string
	Active       bool
}

type RefreshToken struct {
	Identifier  string
	AccessToken string
	Expiry      time.Time
	Revoked     bool
}
