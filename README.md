# README

go-oauth2 provides a thin wrapper of [golang.org/x/oauth2](https://godoc.org/golang.org/x/oauth2).
This currently supports four providers, GitHub, Facebook, Google and Slack.

See [an example](./examples/main.go)

Here is the essence.

Initilization.
```golang
oauth2.WithProviders(
	github.New(oauth2.NewConfig(oauth2.GITHUB, "http://localhost:8080/auth/callback/"+oauth2.GITHUB)),
)
```

Redirecting to provider page.
```golang
loginURL := oauth2.ProviderByName(provider).Config().AuthCodeURL("state")
w.Header().Set("Location", loginURL)
w.WriteHeader(http.StatusTemporaryRedirect)
```

Retrieving profile.
```golang
p := oauth2.ProviderByName(provider)
u, _:= oauth2.ProfileByCode(p, r.FormValue("code"))
log.Println("Name", u.Name())
// use u.Token() ...
```

## License
[MIT License](http://opensource.org/licenses/MIT)
