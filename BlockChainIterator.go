/**
  @author:BOEN
  @data:2022/8/21
  @note:
**/
package main

import (
	"github.com/boltdb/bolt"
	"github.com/go-kit/kit/log"
	"os"
)

// BlockchainIterator 区块链迭代器
type BlockchainIterator struct {
	currentHash []byte
	DB          *bolt.DB
}

/**
 * @Title Iterator
 * @Description //通过区块链实例返回一个与数据库连接的区块链迭代器
 * @Author Cofeesy 18:01 2022/8/21
 * @Param nil
 * @Return *BlockchainIterator
 **/
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.DB}

	return bci
}

/**
 * @Title Next
 * @Description //返回链中的下一个块(从区块链尾部开始迭代)
 * @Author Cofeesy 18:06 2022/8/21
 * @Param nil
 * @Return *Block
 **/
func (bci *BlockchainIterator) Next() *Block {
	logger := log.NewLogfmtLogger(os.Stdout)
	var block *Block

	if err := bci.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(bci.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	}); err != nil {
		logger.Log("bciView", err)
	}

	bci.currentHash = block.PrevBlockHash

	return block
}
