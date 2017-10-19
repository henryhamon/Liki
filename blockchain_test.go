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

/*
func TestAddBlock(t *testing.T) {
	bc := NewBlockchain()

	genesis := bc.blocks[0]

	bc.AddBlock("Send 1 Liki to Jane Doe")

	if len(bc.blocks) != 2 {
		t.Errorf("Blockchain size was incorrect, got: %x, want: %x.", len(bc.blocks), 2)
	}

	if string(bc.blocks[len(bc.blocks)-1].PrevBlockHash) != string(genesis.Hash) {

		t.Errorf("Blockchain was incorrect, got: %x, want: %x.", bc.blocks[len(bc.blocks)-1].PrevBlockHash, genesis.Hash)
	}
}
//*/
