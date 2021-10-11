package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/chorockuin/chorocoin/db"
	"github.com/chorockuin/chorocoin/utils"
)

func (b *Block) make_hash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(b)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	fmt.Printf("AddBlock() NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			fmt.Printf("Blockchain() NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				b.restore(checkpoint)
			}
			fmt.Printf("Blockchain() NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
		})
	}
	return b
}
