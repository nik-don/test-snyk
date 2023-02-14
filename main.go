package main

import (
	"fmt"

	"github.com/ipfs/go-bitfield"
)

func main(){
	// use github.com/ipfs/go-bitfield
	a, _:= go-bitfield.NewBitfield()
	fmt.Println(a)
}