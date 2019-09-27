package ping

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"io"
	"net/http"
)

func Exec(ctx context.Context, url string, contentType string, body io.Reader) (r *http.Response, err error) {
	client := urlfetch.Client(ctx)

	return client.Post(url, contentType, body)
}
