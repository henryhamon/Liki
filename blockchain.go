package main

//Blockchain linked list
type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
