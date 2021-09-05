package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data      string
	hash      string
	prev_hash string
}

func (b *block) make_hash() {
	hash := sha256.Sum256([]byte(b.data + b.prev_hash))
	b.hash = fmt.Sprintf("%x", hash)
}

type blockchain struct {
	blocks []*block
}

func (b *blockchain) get_last_hash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) create_block(data string) *block {
	new_block := block{data, "", b.get_last_hash()}
	new_block.make_hash()
	return &new_block
}

func (b *blockchain) Add_block(data string) {
	b.blocks = append(b.blocks, b.create_block(data))
}

func (b *blockchain) List_blocks() {
	for _, block := range b.blocks {
		fmt.Printf("data:%s\n", block.data)
		fmt.Printf("hash:%s\n", block.hash)
		fmt.Printf("prev_hash:%s\n", block.prev_hash)
	}
}

var b *blockchain
var once sync.Once

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.Add_block("genesis")
		})
	}
	return b
}
