package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"

	"github.com/chorockuin/chorocoin/db"
	"github.com/chorockuin/chorocoin/utils"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

func (b *Block) make_hash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
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
	bc.CurrentDifficulty = block.Difficulty
	fmt.Printf("AddBlock() NewestHash: %s\nHeight: %d\n", bc.NewestHash, bc.Height)
	bc.persist()
}

func (bc *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := bc.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp - lastRecalculatedBlock.Timestamp) / 60
	expectedTime := difficultyInterval * blockInterval
	if actualTime < (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func Blockchain() *blockchain {
	if bc == nil {
		once.Do(func() {
			bc = &blockchain{NewestHash: "", Height: 0}
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
