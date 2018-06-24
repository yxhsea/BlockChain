package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//定义区块
type Block struct {
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一个区块的hash值
	Hash          string //当前区块哈希值

	Data string //区块数据
}

//计算Hash
func calculateHash(b *Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

//生成新的区块
func GenerateNewBlock(preBlock *Block, data string) *Block {
	newBlock := &Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Hash = calculateHash(newBlock)
	newBlock.Data = data
	return newBlock
}

//生成创始区块
func GenerateGenesisBlock() *Block {
	preBlock := &Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
