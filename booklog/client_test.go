package booklog

import "testing"

const (
	host = "http://api.booklog.jp/"
)

func TestNewClient(t *testing.T) {
	cli, err := NewClient("", nil)
	if err == nil {
		t.Fatal("error should be not nil")
	}

	cli, err = NewClient(host, nil)
	if err != nil {
		t.Fatal("error should be nil")
	}
	if cli == nil {
		t.Fatal("error should be not nil")
	}
}
