package clickup

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	if _, err := NewClient(nil); err != nil {
		t.Fatal("Failed to create Client")
	}
}

func TestSetHeader(t *testing.T) {
	client, _ := NewClient(nil)
	client.SetHeader("Header1", "hogehoge")

	if client.headers["Header1"] != "hogehoge" {
		t.Fatal("Header1 is wrong")
	}
}
