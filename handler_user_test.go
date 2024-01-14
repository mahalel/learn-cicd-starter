package main

import (
	"testing"
)

func TestGenerateRandomSHA256Hash(t *testing.T) {
	hash, err := generateRandomSHA256Hash()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(hash) != len("2e2e12b34077247ae90dc2e2e24d9e79dec5f2903270f91025addf32b9a021b0") {
		t.Errorf("length of generateRandomSHA256Hash() = %q, want %q", len(hash), len("2e2e12b34077247ae90dc2e2e24d9e79dec5f2903270f91025addf32b9a021b0"))
	}
}
