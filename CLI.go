/**
  @author:BOEN
  @data:2022/8/21
  @note:
**/
package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"os"
	"strconv"
)

type CLI struct {
	bc *Blockchain
}

/**
 * @Title printUsage
 * @Description //打印命令行交互信息
 * @Author Cofeesy 20:23 2022/8/21
 * @Param nil
 * @Return nil
 **/
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addBlock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printChain - print all the blocks of the blockchain")
}

/**
 * @Title validateArgs
 * @Description //命令行参数验证
 * @Author Cofeesy 20:24 2022/8/21
 * @Param nil
 * @Return nil
 **/
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

/**
 * @Title addBlock
 * @Description //调用底层函数增加区块
 * @Author Cofeesy 18:46 2022/8/21
 * @Param data string
 * @Return nil
 **/
func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

/**
 * @Title printChain
 * @Description //根据迭代对象打印每个区块的信息
 * @Author Cofeesy 19:06 2022/8/21
 * @Param nil
 * @Return nil
 **/
func (cli *CLI) printChain() {
	//返回需要打印的Iterator对象
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.isValid()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

/**
 * @Title Run
 * @Description //解析命令行参数
 * @Author Cofeesy 18:21 2022/8/21
 * @Param nil
 * @Return nil
 **/
func (cli *CLI) Run() {
	logger := log.NewLogfmtLogger(os.Stdout)

	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			logger.Log("addBlockCmd", err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			logger.Log("printChainCmd", err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)

		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
