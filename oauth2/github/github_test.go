package github

import (
	"testing"

	oauth2 "github.com/tmtk75/oauth2"
	xoauth2 "golang.org/x/oauth2"
)

func TestNew(t *testing.T) {
	p := New(&xoauth2.Config{})
	if !(p.Name() == oauth2.GITHUB) {
		t.Fatalf("it should be %v", oauth2.GITHUB)
	}

	if len(p.Config().Scopes) != 0 {
		t.Fatalf("wrong default scopes: %v", defaultScopes)
	}
}
