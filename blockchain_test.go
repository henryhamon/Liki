package main

import (
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain()

	if len(bc.blocks) != 1 {
		t.Errorf("Blockchain was incorrect, got: %x, want: %x.", len(bc.blocks), 1)
	}
}
