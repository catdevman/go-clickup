package clickup

import (
	"net/http"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(nil)
	if err != nil {
		t.Fatal("Failed to create Client")
	}

	if c.baseURL.String() != BASE_URL {
		t.Fatal("Failed to set base URL")
	}
}

func TestSetHeader(t *testing.T) {
	client, _ := NewClient(nil)
	client.SetHeader("Header1", "hogehoge")

	if client.headers["Header1"] != "hogehoge" {
		t.Fatal("Header1 is wrong")
	}
}

func TestSetCredential(t *testing.T) {
	client, _ := NewClient(nil)
	client.SetCredential("pk_test")

	if client.credential != "pk_test" {
		t.Fatal("client.credential returns wrong key: " + client.credential)
	}
}

func TestIncludeHeaders(t *testing.T) {
	client, _ := NewClient(nil)
	client.headers = map[string]string{
		"Header1":      "1",
		"Header2":      "2",
		"Content-Type": "application/json",
	}

	req, _ := http.NewRequest("POST", "localhost", strings.NewReader(""))
	client.includeHeaders(req)

	if len(req.Header) != 3 {
		t.Fatal("req.Header length does not match")
	}

	for k, v := range req.Header {
		switch k {
		case "Header1":
			if v[0] != "1" {
				t.Fatalf(`%s header expect "1", but got "%s"`, k, v[0])
			}
		case "Header2":
			if v[0] != "2" {
				t.Fatalf(`%s header expect "2", but got "%s"`, k, v[0])
			}
		case "Content-Type":
			if v[0] != "application/json" {
				t.Fatalf(`%s header expect "2", but got "%s"`, k, v[0])
			}
		}
	}
}
