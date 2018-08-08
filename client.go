package pester

import (
	"net/http"
	"sync"
	"time"
)

func NewDefaultConfig() Config {
	return defaultConfig
}

// New constructs a new DefaultClient with sensible default values
func New(config Config, withClient *http.Client) Client {
	cl := &client{
		Config:         config,
		ErrLog:         []ErrEntry{},
		wg:             &sync.WaitGroup{},
		RetryOnHTTP429: false,
	}

	if withClient != nil {
		cl.hc = withClient
	}
	return cl
}

type Config struct {
	Concurrency int
	MaxRetries  int
	Backoff     BackoffStrategy
	KeepLog     bool
	LogHook     LogHook
}

// client wraps the http client and exposes all the functionality of the http.client.
// Additionally, client provides pester specific values for handling resiliency.
type client struct {
	// Configuration passed by the user during construction
	Config
	// wrap it to provide access to http built ins
	hc *http.Client

	Transport     http.RoundTripper
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
	Timeout       time.Duration

	// pester specific
	SuccessReqNum   int
	SuccessRetryNum int

	wg *sync.WaitGroup

	sync.Mutex
	ErrLog         []ErrEntry
	RetryOnHTTP429 bool
}
