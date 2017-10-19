package main

import (
	"testing"
	//	"crypto/sha256"
)

func TestCheck(t *testing.T) {
	block := &Block{1508419940, []byte("Just some test data in the string"), []byte(""), []byte{0, 0, 2, 167, 49, 21, 5, 188, 160, 176, 220, 239, 84, 18, 125, 43, 45, 139, 109, 48, 170, 141, 41, 243, 163, 220, 56, 57, 78, 213, 31, 236}, 5382854}

	pow := NewProofOfWork(block)

	if !pow.Check() {
		t.Errorf("Hash was incorrect.")
	}

	/*
		0000023ad068a72b331faadd4889ea8ff9deb73aa8374175fbb667e66704c937
		0000023ad068a72b331faadd4889ea8ff9deb73aa8374175fbb667e66704c937

	*/
}

func TestFindProof(t *testing.T) {
	block := &Block{1508419940, []byte("Just some test data in the string"), []byte(""), []byte(""), 5382000}
	pow := NewProofOfWork(block)

	pow.FindProof()

	var nonce int64 = 5382854
	//proof := []byte{0, 0, 2, 167, 49, 21, 5, 188, 160, 176, 220, 239, 84, 18, 125, 43, 45, 139, 109, 48, 170, 141, 41, 243, 163, 220, 56, 57, 78, 213, 31, 236}

	if block.Nonce != nonce {
		t.Errorf("Nonce was incorrect, got: %x, want: %x.", block.Nonce, nonce)
	}
}
