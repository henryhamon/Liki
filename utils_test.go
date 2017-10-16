package main

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestIntToHex(t *testing.T) {
	buf := new(bytes.Buffer)
	target := int64(256 - 24)
	err := binary.Write(buf, binary.BigEndian, target)
	got := string(IntToHex(target))
	want := string(buf.Bytes())

	if err != nil || got != want {
		t.Errorf("IntToHex was incorrect, got: %x, want: %x.", got, want)
	}
}
