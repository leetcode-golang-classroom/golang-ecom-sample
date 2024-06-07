package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("eror hashing password: %v", err)
	}
	if hash == "" {
		t.Error("expected hash to be not empty")
	}
	if hash == "passowrd" {
		t.Error("expected hash to be different from password")
	}
}
