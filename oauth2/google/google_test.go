package google

import (
	"testing"

	"github.com/tmtk75/go-oauth2/oauth2"
	xoauth2 "golang.org/x/oauth2"
)

func TestNew(t *testing.T) {
	p := New(&xoauth2.Config{})
	if !(p.Name() == oauth2.GOOGLE) {
		t.Fatalf("it should be %v", oauth2.GOOGLE)
	}

	if len(p.Config().Scopes) != 1 {
		t.Fatalf("wrong default scopes: %v", defaultScopes)
	}

	if p.Config().Scopes[0] != defaultScopes[0] {
		t.Fatalf("wrong default scopes: %v", defaultScopes)
	}
}
