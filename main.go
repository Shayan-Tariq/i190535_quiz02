// main.go
package main

import "E:/Blockchain/quiz2/blockchain"

func main() {
	// Create some example blocks
	blocks := []blockchain.Block{
		{Data: "Transaction 1"},
		{Data: "Transaction 2"},
		{Data: "Transaction 3"},
	}

	// Display all blocks
	blockchain.DisplayAllBlocks(blocks)
}
