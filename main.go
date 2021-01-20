package main

import (
	"fmt"
	//"github.com/Ahad-Parmar/Block-chain Golang/blockchain"
	"github.com\Ahad-Parmar\Block-chain Golang\blockchain"

)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("second Block After Genesis")
	chain.AddBlock("third Block After Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
