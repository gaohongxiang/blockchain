package BLC

type Blockchain struct {
	Blocks []*Block //存储有序区块
}

// 创建一个带有创世区块的区块链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
// 添加区块
func(blockchain *Blockchain) AddBlock(data string) {
	preBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(preBlock.Hash, data)

	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}