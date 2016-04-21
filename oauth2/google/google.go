package google

import (
	"github.com/tmtk75/go-oauth2/oauth2"

	xoauth2 "golang.org/x/oauth2"
	xgoogle "golang.org/x/oauth2/google"
)

const profileEndpoint = "https://www.googleapis.com/oauth2/v1/userinfo"

var defaultScopes = []string{"https://www.googleapis.com/auth/userinfo.profile"}

type googleProvider struct {
	config *xoauth2.Config
}

// New returns oauth2.Provider for google
func New(c *xoauth2.Config) oauth2.Provider {
	c.Endpoint = xgoogle.Endpoint
	if c.Scopes == nil {
		c.Scopes = defaultScopes
	}
	return &googleProvider{config: c}
}

func (g *googleProvider) Name() string {
	return oauth2.GOOGLE
}

func (g googleProvider) Config() *xoauth2.Config {
	return g.config
}

func (g googleProvider) Profile(token *xoauth2.Token) (oauth2.Profile, error) {
	data, err := oauth2.GetProfileData(profileEndpoint, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &googleUser{token: token, profile: data}, nil
}

type googleUser struct {
	token   *xoauth2.Token
	profile map[string]interface{}
}

func (u *googleUser) Token() *xoauth2.Token {
	return u.token
}

func (u *googleUser) Name() string {
	return u.profile["name"].(string)
}

func (u *googleUser) Nickname() string {
	return u.profile["nickname"].(string)
}

func (u *googleUser) AvatarURL() string {
	return u.profile["picture"].(string)
}
