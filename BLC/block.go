package BLC

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Timestamp int64			//时间戳,创建区块时的时间
	PrevBlockHash []byte	//父区块哈希
	Data []byte				//交易数据
	Hash []byte				//当前区块哈希
}

// 设置Hash
func(block *Block) SetHash(){
	// 将时间戳转换为字节数组
	// 先转换为二进制字符串，再转换为字节数组
	timestamp := []byte(strconv.FormatInt(block.Timestamp,2))
	// 将除了Hash以外的其他属性以字节数组的形式拼接起来
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp},[]byte{})
	// 将拼接起来的数据进行hash
	hash := sha256.Sum256(headers)
	// 将hash值赋值给Hash属性
	block.Hash = hash[:]
}

// 创建区块方法
func NewBlock(prevBlockHash []byte, data string) *Block {
	// 创建区块
	block := &Block{time.Now().Unix(),prevBlockHash,[]byte(data),[]byte{}}
	// 设置当前区块的hash值
	block.SetHash()
	// 返回区块
	return block
}

//创世区块
func NewGenesisBlock() *Block {
	return NewBlock([]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, "Genenis Block")
}