package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"testing"
)

func TestSetHash(t *testing.T) {
	block := &Block{1257894000, []byte("First Hash"), []byte(""), []byte{}, 0}

	block.SetHash()

	timestamp := []byte(strconv.FormatInt(1257894000, 10))
	headers := bytes.Join([][]byte{[]byte(""), []byte("First Hash"), timestamp}, []byte{})
	shahash := sha256.Sum256(headers)
	hash := shahash[:]

	if string(block.Hash) != string(hash) {
		t.Errorf("Hash was incorrect, got: %x, want: %x.", block.Hash, hash)
	}
}

func TestGenesisBlock(t *testing.T) {

	block := GenesisBlock()
	if string(block.PrevBlockHash) != string([]byte{}) {
		t.Errorf("Genesis Block was incorrect, got: %x, want: %x.", block.PrevBlockHash, []byte{})
	}
}

func TestNew(t *testing.T) {
	genesis := GenesisBlock()
	block := NewBlock("New Block", genesis.Hash)

	if block.Hash == nil || string(block.PrevBlockHash) != string(genesis.Hash) {
		t.Errorf("New Block PrevHash was incorrect, got: %x, want: %x.", block.PrevBlockHash, genesis.Hash)
	}
}
