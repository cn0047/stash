package internal

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/to-com/wp/config"
	"net/http"
)

func NewHTTPClient(cfg *config.Config) *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = cfg.HTTPRetryNum
	retryClient.CheckRetry = retryablehttp.ErrorPropagatedRetryPolicy

	client := retryClient.StandardClient()
	client.Timeout = cfg.HTTPRequestTimeout

	return client
}
