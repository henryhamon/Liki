package main

import (
	"testing"
)

func init() {

}

func TestTransactionID(t *testing.T) {
	transaction := NewCoinbaseTX("XYZ", "The Times 03/Set/2013 Ylvis asked What does the Fox says?")
	if SHAString(transaction.ID) != "58ec9d4930a33909096869892b5d6601fda3de1a8397da603f0cb31909f109f8" {
		t.Errorf("Transaction ID was incorrect.")
	}
}

func TestNewCoinBase(t *testing.T) {
	transaction := NewCoinbaseTX("XYZ", "The Times 03/Set/2013 Ylvis asked What does the Fox says?")
	if IsNil(transaction) {
		t.Errorf("Transaction was incorrect.")
	}
}
