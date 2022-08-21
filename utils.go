/**
  @author:BOEN
  @data:2022/8/21
  @note:
**/
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
)

/**
 * @Title IntToHex
 * @Description //将Int64转换为字节数组
 * @Author Cofeesy 18:56 2022/8/13
 * @Param num int64
 * @Return []byte
 **/
func IntToHex(num int64) []byte {
	//创建一个缓冲
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

/**
 * @Title Serialize
 * @Description //将一个区块序列化为字符数组-->用于存储
 * @Author Cofeesy 16:10 2022/8/15
 * @Param nil
 * @Return []byte
 **/
func (block *Block) SerializeBlock() []byte {

	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

/**
 * @Title DeserializeBlock
 * @Description //将一个区块反序列化-->将区块信息取出来
 * @Author Cofeesy 16:14 2022/8/15
 * @Param
 * @Return
 **/
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
