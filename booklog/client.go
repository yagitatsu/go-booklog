package booklog

import (
	"fmt"
	"net/http"
	"net/url"

	"encoding/json"
	"github.com/pkg/errors"
)

const (
	Status_All int = iota
	Status_WantToRead
	Status_Reading
	Status_Done
	Status_Tsundoku
)

type (
	// Client for booklog
	Client struct {
		url    *url.URL
		client *http.Client
	}

	// GetOptions is parameters for Get method.
	GetOptions struct {
		Count  int // items limit. default value is 5.
		Status int // 0 is all,

	}
	GetResult struct {
		Shelf struct {
			ID       string `json:"id"`
			Account  string `json:"acount"`
			Name     string `json:"name"`
			ImageURL string `json:"image_url"`
		} `json:"tana"`
		Categories []string `json:"category"`
		Books      []Book   `json:"books"`
	}

	Book struct {
		ID      string `json:"id"`
		ASIN    string `json:"asin"`
		URL     string `json:"url"`
		Title   string `json:"title"`
		Author  string `json:"author"`
		Image   string `json:"image"`
		Width   string `json:"width"`
		Height  string `json:"height"`
		Catalog string `json:"catalog"`
	}
)

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
		url:    u,
		client: cli,
	}, nil
}

// Get returns user's booklogs.
func (c *Client) Get(id string, opts *GetOptions) (GetResult, error) {
	const errtag = "client.Get failed"
	u := *c.url
	u.Path = u.Path + "/" + id
	if opts != nil {
		q := u.Query()
		if opts.Count > 0 {
			q.Set("count", fmt.Sprint(opts.Count))
		}
		if opts.Status > 0 {
			q.Set("status", fmt.Sprint(opts.Status))
		}
		u.RawQuery = q.Encode()
	}
	resp, err := c.client.Get(u.String())
	if err != nil {
		return GetResult{}, errors.Wrap(err, errtag)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return GetResult{}, fmt.Errorf("%s: status code = %d", errtag, resp.StatusCode)
	}
	var r GetResult
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return GetResult{}, err
	}

	return r, nil
}
