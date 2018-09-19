package core

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

var (
	privateRootPrefix    = []byte("Priv")
	privateReceiptPrefix = []byte("PrivR")
)

// GetPrivateStateRoot gets the root(hash) for private state associated with the root of Merkle tree in public chain.
func GetPrivateStateRoot(db ethdb.Database, blockRoot common.Hash) common.Hash {
	root, _ := db.Get(append(privateRootPrefix, blockRoot[:]...))
	return common.BytesToHash(root)
}

// WritePrivateStateRoot writes the root(hash) for private state associated with the root of Merkle tree in public chain.
func WritePrivateStateRoot(db ethdb.Database, blockRoot, root common.Hash) error {
	return db.Put(append(privateRootPrefix, blockRoot[:]...), root[:])
}

// WritePrivateReceipt writes private receipt associated with specified transaction.
func WritePrivateReceipt(receipt *types.Receipt, txHash common.Hash, db *trie.Database) error {
	// Generate hash combining tx hash and private receipt prefix.
	// It aims at avoiding conflict.
	contentToHash := bytes.Join([][]byte{
		privateReceiptPrefix,
		txHash.Bytes(),
	}, []byte{})
	hasher := sha3.NewKeccak256()
	hasher.Write(contentToHash)
	hashBytes := hasher.Sum(nil)
	hash := common.BytesToHash(hashBytes)

	// Write receipt to trie db.
	storageReceipt := (*types.ReceiptForStorage)(receipt)
	bytesToWrite, _ := rlp.EncodeToBytes(storageReceipt)
	db.InsertBlob(hash, bytesToWrite)
	log.Info("Write private transaction receipt", "hash", hash, "content", bytesToWrite)
	return nil
}
