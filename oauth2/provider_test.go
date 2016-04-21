package oauth2

import (
	"math/rand"
	"testing"

	xoauth2 "golang.org/x/oauth2"
)

type mockProvider struct {
	name string
}

func (m mockProvider) Name() string {
	return m.name
}

func (m mockProvider) Config() *xoauth2.Config {
	return nil
}

func (m mockProvider) Profile(token *xoauth2.Token) (Profile, error) {
	return nil, nil
}

func TestProviderByName(t *testing.T) {
	foo := mockProvider{name: "foo"}
	WithProviders(foo)
	p := ProviderByName("foo")
	if !(p == foo) {
		t.Fatalf("should return %v", foo)
	}
}

func TestProfile(t *testing.T) {
	foo := &mockProvider{name: "foo"}
	var token *xoauth2.Token
	_, err := foo.Profile(token)
	if !(err == nil) {
		t.Fatalf("%v", err)
	}
}

func TestNewConfig(t *testing.T) {
	randStringBytes := func(n int) string {
		const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		b := make([]byte, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	defer func() {
		if r := recover(); r != nil {
			// Expect panic
		} else {
			t.Fatalf("NewConfig panic is expected")
		}
	}()

	s := randStringBytes(8)
	NewConfig(s, "http://callback-url/"+s)
}
