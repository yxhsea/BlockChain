package core

import (
	"fmt"
	"log"
)

//定义区块链
type BlockChain struct {
	Blocks []*Block
}

//创建一个区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	blockChain := &BlockChain{}
	blockChain.AppendBlock(genesisBlock)
	return blockChain
}

//记录区块数据
func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(preBlock, data)
	bc.AppendBlock(newBlock)
}

//往区块链添加区块
func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(newBlock, bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
	return
}

//输出区块链信息
func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index : %d\n", block.Index)
		fmt.Printf("Prev.Hash : %s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash : %s\n", block.Hash)
		fmt.Printf("Curr.Data : %s\n", block.Data)
		fmt.Printf("Curr.Timestamp : %d\n", block.Timestamp)
		fmt.Println("==========================================")
	}
}

//验证区块
func isValid(newBlock *Block, oldBlock *Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
