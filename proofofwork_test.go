package main

import (
	"bytes"
	"testing"
)

func TestNewProofOfWork(t *testing.T) {

	block := NewBlock("New Block", []byte{})
	pow := NewProofOfWork(block)

	if pow.block != block {
		t.Errorf("PoW was incorrect, got: %x, want: %x.", pow.block, block)
	}
}

func TestPrepareData(t *testing.T) {
	block := NewBlock("New Block", []byte{})
	pow := NewProofOfWork(block)

	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(256 - 24)),
			IntToHex(int64(2)),
		},
		[]byte{},
	)

	powdata := pow.prepareData(2)
	if string(powdata) != string(data) {
		t.Errorf("PoW Data was incorrect, got: %x, want: %x.", powdata, data)
	}
}
