package main

import (
	"fmt"
	"testing"

	zkt "github.com/scroll-tech/zktrie/types"
)

func TestZKT(t *testing.T) {
	t.Run("test zkt package", func(t *testing.T) {
		k1 := zkt.Byte32{1}
		h, _ := k1.Hash()
		fmt.Println("The hash of k1: ", h)
	})
}

// func NewTesingMT(t *testing.T, n int) {

// }

func TestMT_ForgeProof(t *testing.T) {
	zkTrie := NewTesingMT(t, 10)

}
