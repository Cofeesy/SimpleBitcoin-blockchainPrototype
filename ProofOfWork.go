/**
  @author:BOEN
  @data:2022/8/21
  @note:
**/
package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

//定义挖矿的难度值
const targetBit = 2

//定义最大nonce值，防止溢出
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	//即将生成的区块对象
	Block *Block
	//生成区块的难度
	target *big.Int
}

/**
 * @Title prepareData
 * @Description //1.准备数据：拼接除了区块里面除了Hash以外的其他所有数据
 * @Author Cofeesy 22:13 2022/8/14
 * @Param nonce int
 * @Return []byte
 **/
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.TimeStamp),
			IntToHex(int64(targetBit)),
			IntToHex(nonce),
		},
		[]byte{},
	)
	return data
}

/**
 * @Title isValid
 * @Description //对工作量证明进行验证
 * @Author Cofeesy 0:02 2022/8/15
 * @Param
 * @Return
 **/
func (pow *ProofOfWork) isValid() bool {
	var hashInt big.Int
	//生成Hash

	data := pow.prepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	//将hash存储到hashInt中
	hashInt.SetBytes(hash[:])

	//判断hash
	// Cmp compares x and y and returns:
	//
	//   -1 if x <  y
	//    0 if x == y
	//   +1 if x >  y
	//
	isValid := pow.target.Cmp(&hashInt) == 1

	return isValid
}

/**
 * @Title Run
 * @Description //2.用SHA-256对数据进行Hash; 3.将Hash转换成一个大整数;4.将这个大整数与目标进行比较
 * @Author Cofeesy 16:21 2022/8/14
 * @Param
 * @Return
 **/
func (pow *ProofOfWork) Run() ([]byte, int64) {
	var hash [32]byte
	var hashInt big.Int //存储新生成的区块Hash

	//先初始化nonce为0
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.Block.Data)

	//通过nonce值碰撞产生符合逻辑的target值
	for nonce < maxNonce {
		//调用拼接函数
		dataBytes := pow.prepareData(int64(nonce))

		//生成Hash
		hash = sha256.Sum256(dataBytes)

		//将hash存储到hashInt中
		hashInt.SetBytes(hash[:])

		//判断hash
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		if pow.target.Cmp(&hashInt) == 1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return hash[:], int64(nonce)
}

/**
 * @Title NewProofOfWork
 * @Description 创建新的工作量证明对象
 * @Author Cofeesy 16:20 2022/8/14
 * @Param
 * @Return
 **/
func NewProofOfWork(block *Block) *ProofOfWork {
	//创建一个初始值为1的target
	target := big.NewInt(1)

	//左移256-targetBit
	target = target.Lsh(target, 256-targetBit)

	pow := &ProofOfWork{block, target}
	return pow
}
