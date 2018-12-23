package main

import (
	"publicChain/BLC"
	"time"
	"fmt"
)

func main()  {1

	blockchain := BLC.NewBlockchain()
	blockchain.AddBlock("Send 1 BTC To xiaoming From zhongbencong")
	blockchain.AddBlock("Send 2 BTC To xiaohong From zhongbencong")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Data:%s \n", string(block.Data))
		fmt.Printf("Timestamp:%v \n", time.Unix(block.Timestamp,0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("PrevBlockHash:%x \n", block.PrevBlockHash)
		fmt.Printf("Hash:%x \n", block.Hash)

		fmt.Println()

	}

}