package tests

import (
	"testing"

	"interview-copilot/backend/internal/auth"
)

func TestJWT(t *testing.T) {
	secret := "test"
	token, err := auth.Generate(42, secret)
	if err != nil {
		t.Fatal(err)
	}

	id, err := auth.Parse(token, secret)
	if err != nil {
		t.Fatal(err)
	}

	if id != 42 {
		t.Fatalf("expected 42, got %d", id)
	}
}
