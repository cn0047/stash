package go_app

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/user")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}

	url, _ := user.LogoutURL(ctx, "/user")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

// userOAuthHandler
// ‼️ Read more:
// @link: https://cloud.google.com/appengine/docs/standard/go/users/#Go_OAuth_in_Go
func userOAuthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u, err := user.CurrentOAuth(ctx, "")
	if err != nil {
		http.Error(w, "OAuth Authorization header required", http.StatusUnauthorized)
		return
	}

	if !u.Admin {
		http.Error(w, "Admin login only", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, `Welcome, admin user %s!`, u)
}
