// main.go

package main

import "fmt"

type Block struct {
	Data string
}

// DisplayAllBlocks prints information about all blocks in the given slice.
func DisplayAllBlocks(blocks []Block) {
	for i := 0; i < len(blocks); i++ {
		fmt.Printf("Block %d: %s\n", i+1, blocks[i].Data)
	}
}

// NewBlock creates a new block with the specified data.
func NewBlock(data string) Block {
	return Block{
		Data: data,
	}
}
func ModifyBlock(block *Block, data string) {
	block.Data = data
}

func main() {
	// Create some blocks
	block1 := NewBlock("Embrace Masculinty ")
	block2 := NewBlock("Reject Modernity")
	block3 := NewBlock("Thanks")
	ModifyBlock(&block3, "Remember")
	// Display all blocks
	blocks := []Block{block1, block2, block3}
	DisplayAllBlocks(blocks)
}