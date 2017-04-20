package booklog

import (
	"net/http"
	"net/url"
	"github.com/pkg/errors"
)

// Client for booklog
type Client struct {
	url    *url.URL
	client *http.Client
}

// NewClient creates a client for booklog.
func NewClient(host string, cli *http.Client) (*Client, error) {
	const errtag = "booklob.NewClient failed"
	u, err := url.ParseRequestURI(host)
	if err != nil {
		return nil, errors.Wrap(err, errtag)
	}
	u.Path = "/json"
	if cli == nil {
		cli = http.DefaultClient
	}
	return &Client{
		url: u,
		client: cli,
	}, nil
}
