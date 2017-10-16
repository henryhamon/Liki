package main

import (
	"testing"
)

func TestNewProofOfWork(t *testing.T) {

	block := NewBlock("New Block", []byte{})
	pow := NewProofOfWork(block)

	if pow.block != block {
		t.Errorf("Pow was incorrect, got: %x, want: %x.", pow.block, block)
	}
}
