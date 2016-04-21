package facebook

import (
	"github.com/tmtk75/go-oauth2/oauth2"

	xoauth2 "golang.org/x/oauth2"
	xfacebook "golang.org/x/oauth2/facebook"
)

const profileEndpoint = "https://graph.facebook.com/me?fields=email,first_name,last_name,link,about,id,name,picture,location"

var defaultScopes = []string{"email"}

type facebookProvider struct {
	config *xoauth2.Config
}

// New returns oauth2.Provider for facebook
func New(c *xoauth2.Config) oauth2.Provider {
	c.Endpoint = xfacebook.Endpoint
	if c.Scopes == nil {
		c.Scopes = defaultScopes
	}
	return facebookProvider{config: c}
}

func (f facebookProvider) Name() string {
	return oauth2.FACEBOOK
}

func (f facebookProvider) Config() *xoauth2.Config {
	return f.config
}

func (f facebookProvider) Profile(token *xoauth2.Token) (oauth2.Profile, error) {
	data, err := oauth2.GetProfileData(profileEndpoint, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &facebookUser{token: token, profile: data}, nil
}

type facebookUser struct {
	token   *xoauth2.Token
	profile map[string]interface{}
}

func (u *facebookUser) Token() *xoauth2.Token {
	return u.token
}

func (u *facebookUser) Name() string {
	return u.profile["name"].(string)
}

func (u *facebookUser) Nickname() string {
	return u.profile["username"].(string)
}

func (u *facebookUser) AvatarURL() string {
	if a, ok := u.profile["picture"].(map[string]interface{}); ok {
		if b, ok := a["data"].(map[string]interface{}); ok {
			return b["url"].(string)
		}
	}
	return ""
}
