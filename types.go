package pester

import (
	"net/http"
	"io"
	"net/url"
)

// Wait blocks until all pester requests have returned
// Probably not that useful outside of testing.
type Client interface{
	Wait()
	EmbedHTTPClient(hc *http.Client)
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
	SetRetryOnHTTP429(flag bool)

	// Log Functions
	LogString() string
	FormatError(e ErrEntry) string
	LogErrCount() int

	// Get Config, allows to look at, but not change, the config
	GetConfig() Config
}