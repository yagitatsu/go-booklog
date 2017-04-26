package booklog

import "testing"

const (
	host = "http://api.booklog.jp/"
	id   = "tatsuyayagi"
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

func TestClient_Get(t *testing.T) {
	cli, err := NewClient(host, nil)
	if err != nil {
		t.Fatal("error should be nil")
	}
	r, err := cli.Get(id, nil)
	if err != nil {
		t.Fatal("error should be nil")
	}
	t.Logf("result=%v", r)
	if len(r.Books) == 0 {
		t.Fatal("There should be some books")
	}
}
