package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 22

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	pow := &ProofOfWork{b, targetBits}

	return pow
}

func (pow ProofOfWork) ZeroCount() int {
	digest := sha256.Sum256(pow.PrepareHash())
	digestHex := new(big.Int).SetBytes(digest[:])
	return ((sha256.Size * 8) - digestHex.BitLen())
}

func (pow ProofOfWork) PrepareHash() []byte {
	var buf bytes.Buffer

	buf.Write(pow.block.PrevBlockHash)
	buf.Write(pow.block.Data)

	binary.Write(&buf, binary.BigEndian, pow.block.TimeStamp)
	binary.Write(&buf, binary.BigEndian, pow.block.Nonce)

	return buf.Bytes()
}

func (pow ProofOfWork) Check() bool {
	if pow.ZeroCount() >= pow.target {
		return true
	}
	return false
}

// FindProof performs a proof-of-work
func (pow *ProofOfWork) FindProof() {
	var hash [32]byte
	fmt.Printf("\rMining the block containing \"%s\"\n", pow.block.Data)
	for {
		hash = sha256.Sum256(pow.PrepareHash())
		if pow.Check() {
			fmt.Printf("\r%x (yay!)", hash)

			return
		}
		if pow.block.Nonce >= math.MaxInt64 {
			fmt.Println("\n:(")
			return
		}

		pow.block.Nonce++
	}
}
