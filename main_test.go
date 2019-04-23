package main


import (
	"fmt"
	"testing"
)

func TestIdenticon(t* testing.T){
	bytes := []byte{1,2,3}
	fmt.Printf("Hash: %v\n", bytes)
	res :=  Render(bytes)
	fmt.Printf("Result identicon: %v\n", res)
}
