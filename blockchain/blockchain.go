package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data      string
	Hash      string
	Prev_hash string
}

func (b *Block) make_hash() {
	hash := sha256.Sum256([]byte(b.Data + b.Prev_hash))
	b.Hash = fmt.Sprintf("%x", hash)
}

type blockchain struct {
	blocks []*Block
}

func (b *blockchain) get_last_hash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].Hash
	}
	return ""
}

func (b *blockchain) create_block(data string) *Block {
	new_block := Block{data, "", b.get_last_hash()}
	new_block.make_hash()
	return &new_block
}

func (b *blockchain) Add_block(data string) {
	b.blocks = append(b.blocks, b.create_block(data))
}

func (b *blockchain) List_blocks() {
	for _, block := range b.blocks {
		fmt.Printf("data:%s\n", block.Data)
		fmt.Printf("hash:%s\n", block.Hash)
		fmt.Printf("prev_hash:%s\n", block.Prev_hash)
	}
}

var b *blockchain
var once sync.Once

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.Add_block("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
