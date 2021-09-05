package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data      string
	hash      string
	prev_hash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) get_last_hash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) add_block(data string) {
	new_block := block{data, "", ""}
	new_block.prev_hash = b.get_last_hash()
	hash := sha256.Sum256([]byte(new_block.data + new_block.prev_hash))
	new_block.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, new_block)
}

func (b *blockchain) list_blocks() {
	for _, block := range b.blocks {
		fmt.Printf("data:%s\n", block.data)
		fmt.Printf("hash:%s\n", block.hash)
		fmt.Printf("prev_hash:%s\n", block.prev_hash)
	}
}

func main() {
	chain := blockchain{}
	chain.add_block("genesis_block")
	chain.add_block("second_block")
	chain.add_block("third_block")
	chain.list_blocks()
}
