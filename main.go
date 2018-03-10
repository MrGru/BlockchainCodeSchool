package main

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Gru")
	bc.AddBlock("Send 2 more BTC to Gru")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: \n%x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: \n%x\n", block.Hash)
		fmt.Println()
	}
}
