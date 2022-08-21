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

//数据库名字
const dbFile = "blockchain.db"

//数据库表名
const blocksBucket = "blocks"

type Blockchain struct {
	//是区块链里面的最后一个区块的Hash
	Tip []byte
	//Bolt持久化操作对象
	DB *bolt.DB
}

/**
 * @Title NewBlockChain
 * @Description //创建一个带有创世区块的区块链(如果没有区块链：1.创建创世区块-->2.存储到数据库-->3.将创世区块哈希保存为最后一个块的Hash-->4.创建一个新的BlockChain实例，初始化时tip指向创世区块)
 * @Author Cofeesy 11:13 2022/8/14
 * @Param nil
 * @Return *Blockchain
 **/
func NewBlockChain() *Blockchain {

	logger := log.NewLogfmtLogger(os.Stdout)
	var tip []byte
	//打开bolt文件
	//如果该文件不存在，则创建以dbFile为名字的".db"文件
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		logger.Log("open", err)
	}
	//fmt.Println("bolt数据库连接成功")

	//BoltDB文件的两种类型事务操作之一 update-->读写事务
	err = db.Update(func(tx *bolt.Tx) error {

		//获取存储区块的bucket
		//Bucket retrieves a bucket by name. Returns nil if the bucket does not exist
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesisBlock := CreateGenesisBlock("Genesis Block")
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				logger.Log("newCreateBucket", err)
			} else {
				//func (b *Bucket) Put(key []byte, value []byte) error
				//这里以需要存入的区块的Hash为key,整个区块序列化后的字节数组作为value存入bucket中
				err = b.Put(genesisBlock.Hash, genesisBlock.SerializeBlock())
				if err != nil {
					logger.Log("newPutBlockHash", err)
				} else {
					//然后将此存入的区块作为最新的区块.以l为key,该区块Hash为value存入bucket中
					err = b.Put([]byte("l"), genesisBlock.Hash)
					if err != nil {
						logger.Log("newPutLHash", err)
					} else {
						//更新tip
						tip = genesisBlock.Hash
					}
				}
			}
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	bc := &Blockchain{tip, db}
	return bc

}

/**
 * @Title AddBlock
 * @Description //在创建好的区块链中添加区块
 * @Author Cofeesy 11:42 2022/8/14
 * @Param height, int64, data string, prevHash []byte
 * @Return
 **/
func (bc *Blockchain) AddBlock(data string) {
	logger := log.NewLogfmtLogger(os.Stdout)
	//用来接收最后一个区块的Hash值
	var lastHash []byte

	//BoltDB文件的两种类型事务操作之一 View:只读业务
	if err := bc.DB.View(func(tx *bolt.Tx) error {
		//获取bucket
		b := tx.Bucket([]byte(blocksBucket))
		//获取bucket中的l键对应的值
		lastHash = b.Get([]byte("l"))

		return nil
	}); err != nil {
		logger.Log("view", err)
	}
	newBlock := NewBlock(lastHash, data)

	if err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if errPut := b.Put(newBlock.Hash, newBlock.SerializeBlock()); errPut != nil {
			logger.Log("addPutBlockHash", errPut)
		} else {
			if errPut = b.Put([]byte("l"), newBlock.Hash); errPut != nil {
				logger.Log("addPutLHash", errPut)
			} else {
				bc.Tip = newBlock.Hash
			}
		}
		return nil
	}); err != nil {
		logger.Log("addUpdate", err)
	}

}
