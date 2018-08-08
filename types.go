package pester

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Wait blocks until all pester requests have returned
// Probably not that useful outside of testing.
type Client interface {
	Wait()
	EmbedHTTPClient(hc *http.Client)
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
	SetRetryOnHTTP429(flag bool)

	// Log Functions
	GetErrLog() []ErrEntry
	LogString() string
	FormatError(e ErrEntry) string
	LogErrCount() int

	// Get Config, allows to look at, but not change, the config
	GetConfig() Config
}

//ErrUnexpectedMethod occurs when an http.client method is unable to be mapped from a calling method in the pester client
var ErrUnexpectedMethod = errors.New("unexpected client method, must be one of Do, Get, Head, Post, or PostFrom")

// ErrReadingBody happens when we cannot read the body bytes
var ErrReadingBody = errors.New("error reading body")

// ErrReadingRequestBody happens when we cannot read the request body bytes
var ErrReadingRequestBody = errors.New("error reading request body")

// ErrEntry is used to provide the LogString() data and is populated
// each time an error happens if KeepLog is set.
// ErrEntry.Retry is deprecated in favor of ErrEntry.Attempt
type ErrEntry struct {
	Time    time.Time
	Method  string
	URL     string
	Verb    string
	Request int
	Retry   int
	Attempt int
	Err     error
}

// LogHook is used to log attempts as they happen. This function is never called,
// however, if KeepLog is set to true.
type LogHook func(e ErrEntry)
