package main

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBase58Decode(t *testing.T) {
	rawHash := "00010966776006953D5567439E5E39F86A0D273BEED61967F6"
	hash, err := hex.DecodeString(rawHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("hash %x\n", hash)
	encoded := Base58Encode(hash)
	assert.Equal(t, "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM", string(encoded))
}
