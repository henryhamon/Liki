//
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int64
}

func NewBlock(data string, prevBlock []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlock, []byte{}, 0}

	// block.SetHash()
	pow := NewProofOfWork(block)

	pow.FindProof()
	hash := sha256.Sum256(pow.PrepareHash())
	block.Hash = hash[:]

	return block
}

// Calculate and Set Block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func GenesisBlock() *Block {
	block := &Block{time.Now().Unix(), []byte("Genesis Block"), []byte(""), []byte{}, 0}

	block.SetHash()
	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}
