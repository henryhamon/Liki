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

func TestSerialize(t *testing.T) {
	block := &Block{1508419940, []byte("Just some test data in the string"), []byte(""), []byte{0, 0, 2, 167, 49, 21, 5, 188, 160, 176, 220, 239, 84, 18, 125, 43, 45, 139, 109, 48, 170, 141, 41, 243, 163, 220, 56, 57, 78, 213, 31, 236}, 5382854}

	proof := []byte{79, 255, 129, 3, 1, 1, 5, 66, 108, 111, 99, 107, 1, 255, 130, 0, 1, 5, 1, 9, 84, 105, 109, 101, 83, 116, 97, 109, 112, 1, 4, 0, 1, 4, 68, 97, 116, 97, 1, 10, 0, 1, 13, 80, 114, 101, 118, 66, 108, 111, 99, 107, 72, 97, 115, 104, 1, 10, 0, 1, 4, 72, 97, 115, 104, 1, 10, 0, 1, 5, 78, 111, 110, 99, 101, 1, 4, 0, 0, 0, 83, 255, 130, 1, 252, 179, 209, 82, 200, 1, 33, 74, 117, 115, 116, 32, 115, 111, 109, 101, 32, 116, 101, 115, 116, 32, 100, 97, 116, 97, 32, 105, 110, 32, 116, 104, 101, 32, 115, 116, 114, 105, 110, 103, 2, 32, 0, 0, 2, 167, 49, 21, 5, 188, 160, 176, 220, 239, 84, 18, 125, 43, 45, 139, 109, 48, 170, 141, 41, 243, 163, 220, 56, 57, 78, 213, 31, 236, 1, 253, 164, 69, 140, 0}

	if string(block.Serialize()) != string(proof) {
		t.Errorf("Block Serialization was incorrect, got: %x, want: %x.", block.Serialize(), proof)
	}
}

//*/
