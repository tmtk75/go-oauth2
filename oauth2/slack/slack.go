package slack

import (
	"github.com/tmtk75/go-oauth2/oauth2"

	xoauth2 "golang.org/x/oauth2"
	xslack "golang.org/x/oauth2/slack"
)

const profileEndpoint = "https://slack.com/api/auth.test"

var defaultScopes = []string{"read"}

type slackProvider struct {
	config *xoauth2.Config
}

// New returns oauth2.Provider for slack
func New(c *xoauth2.Config) oauth2.Provider {
	c.Endpoint = xslack.Endpoint
	if c.Scopes == nil {
		c.Scopes = defaultScopes
	}
	return slackProvider{config: c}
}

func (s slackProvider) Name() string {
	return oauth2.SLACK
}

func (s slackProvider) Config() *xoauth2.Config {
	return s.config
}

func (s slackProvider) Profile(token *xoauth2.Token) (oauth2.Profile, error) {
	data, err := GetProfileData(token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &slackUser{token: token, profile: data}, nil
}

type slackUser struct {
	token   *xoauth2.Token
	profile map[string]interface{}
}

func (u *slackUser) Token() *xoauth2.Token {
	return u.token
}

func (u *slackUser) Name() string {
	return u.profile["user"].(map[string]interface{})["real_name"].(string)
}

func (u *slackUser) Nickname() string {
	return u.profile["user"].(map[string]interface{})["name"].(string)
}

func (u *slackUser) AvatarURL() string {
	a := u.profile["user"].(map[string]interface{})["profile"]
	return a.(map[string]interface{})["image_72"].(string)
}
