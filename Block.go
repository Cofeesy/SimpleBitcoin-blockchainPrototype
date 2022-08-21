/**
  @author:BOEN
  @data:2022/8/21
  @note:
**/
/**
  @author:BOEN
  @data:2022/8/13
  @note:
**/
package main

import (
	"time"
)

// Block 区块链原型
type Block struct {
	//区块高度
	//Height int64
	//上一个区块的Hash(父Hash)
	PrevBlockHash []byte
	//区块数据
	Data []byte
	//时间戳
	TimeStamp int64
	//本区块Hash
	Hash []byte
	//随机数
	Nonce int64
}

/**
 * @Title SetHash
 * @Description //简易的生成区块Hash
 * @Author Cofeesy 16:50 2022/8/13
 * @Param nil
 * @Return nil
 **/
/*func (block *Block) SetHash() {
	//1.通过自定义函数将区块高度转换为字节数组
	heightBytes := IntToHex(block.Height)
	//fmt.Println(heightBytes)

	//2.将时间戳转换为字符串
	//base:2~36位
	timestring := strconv.FormatInt(block.TimeStamp, 2)
	//fmt.Println(timestring)

	//将时间戳的字符串转换为字节数组
	timestamp := []byte(timestring)

	//3.拼接区块内除了Hash以为的其他值
	headers := bytes.Join([][]byte{
                heightBytes,
                block.PrevBlockHash,
                block.Data,
                timestamp},
           []byte{},
       )

	//4.生成Hash
	hash := sha256.Sum256(headers)

	//上一步返回的是32位的字节数组
	block.Hash = hash[:]
}*/

/**
 * @Title NewBlock
 * @Description //创建一个新的区块
 * @Author Cofeesy 16:50 2022/8/13
 * @Param height int64, prevBlockHash []byte, data string
 * @Return *Block
 **/
func NewBlock(prevBlockHash []byte, data string) *Block {

	//创建一个区块实例(区块Hash初始化为nil)
	block := &Block{prevBlockHash, []byte(data), time.Now().Unix(), nil, 0}

	//设置创建区块实例的Hash
	//block.SetHash()

	//实例化一个ProofOfWork对象
	pow := NewProofOfWork(block)

	//挖矿验证
	hash, nonce := pow.Run()

	//将验证正确后的Hash和Nonce值赋予新生成的区块
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

/**
 * @Title CreateGenesisBlock
 * @Description //创建创世区块
 * @Author Cofeesy 0:19 2022/8/14
 * @Param data string
 * @Return *Block
 **/
func CreateGenesisBlock(data string) *Block {
	return NewBlock([]byte{}, data)
}
