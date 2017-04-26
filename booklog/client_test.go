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
	t.Logf("result=%v", r)
	if err != nil {
		t.Fatal("error should be nil")
	}
	if len(r.Books) == 0 {
		t.Fatal("There should be some books")
	}

	r2, err := cli.Get(id, &GetOptions{
		Count:  1,
		Status: Status_Done,
	})
	if err != nil {
		t.Fatal("error should be nil")
	}

	if len(r2.Books) != 1 {
		t.Fatal("books count should be 1")
	}
}
