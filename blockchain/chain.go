package blockchain

import (
	"crypto/sha256"
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

var bc *blockchain
var once sync.Once

func (bc *blockchain) restore(data []byte) {
	utils.FromBytes(bc, data)
}

func (bc *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(bc))
}

func (bc *blockchain) AddBlock(data string) {
	block := createBlock(data, bc.NewestHash, bc.Height+1)
	bc.NewestHash = block.Hash
	bc.Height = block.Height
	fmt.Printf("AddBlock() NewestHash: %s\nHeight: %d\n", bc.NewestHash, bc.Height)
	bc.persist()
}

func Blockchain() *blockchain {
	if bc == nil {
		once.Do(func() {
			bc = &blockchain{"", 0}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				bc.AddBlock("Genesis")
			} else {
				bc.restore(checkpoint)
			}
		})
	}
	fmt.Println(bc.NewestHash)
	return bc
}
