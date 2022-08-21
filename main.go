/**
  @author:BOEN
  @data:2022/8/21
  @note:
**/
package main

func main() {
	/*
		genesisBlock := BLC.CreateBlockchainGenesisBlock()
		//打印存储创世区块的区块链的地址
		fmt.Println(genesisBlock)

		//打印区块链中的区块数组
		fmt.Println(genesisBlock.Blocks)

		//打印区块数组中的第一个区块-->创世区块的信息
		fmt.Println(genesisBlock.Blocks[0])
	*/

	/*//创建一个区块链
	blockchain := BLC.CreateBlockchainGenesisBlock()

	//添加区块
	blockchain.AddBlock(blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, "李四 Send 50BTC to 张三", blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlock(blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, "王二麻子 Send 100BTC to 李四", blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlock(blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, "张三 Send 10BTC to 王二麻子", blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	fmt.Println(blockchain.Blocks[1].Data)
	fmt.Println(blockchain.Blocks)*/

	bc := NewBlockChain()
	defer bc.DB.Close()

	cli := CLI{bc}
	cli.Run()
}
