package core

import (
	"encoding/gob"
	"fmt"
	"math/rand"

	"github.com/mehdi124/blockcherry/crypto"
	"github.com/mehdi124/blockcherry/types"
)

type TxType byte

const (
	TxTypeCollection = iota //0x0
	TxTypeMint              //0x1
)

type CollectionTx struct {
	Fee      uint64
	MetaData []byte
}

type MintTx struct {
	Fee            uint64
	NFT            types.Hash
	Collection     types.Hash
	MetaData       []byte
	CollectionOwer crypto.PublicKey
	Signature
}

type Transaction struct {
	// Only used for native NFT logic
	TxInner any
	// Any arbitrary data for the VM
	Data      []byte
	To        crypto.PublicKey
	Value     uint64
	From      crypto.PublicKey
	Signature *crypto.Signature
	Nonce     int64

	// cached version of tx data hash
	hash types.Hash
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data:  data,
		Nonce: rand.Int63n(1000000000000000),
	}
}

func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {

	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}

	return tx.hash

	return hasher.Hash(tx)
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {

	hash := tx.Hash(TXHasher{})
	sig, err := privKey.Sign(hash.ToSlice())
	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sig
	return nil
}

func (tx *Transaction) Verify() error {

	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	hash := tx.Hash(TXHasher{})

	if !tx.Signature.Verify(tx.From, hash.ToSlice()) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}

func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}

func init() {
	gob.Register(CollectionTx{})
	gob.Register(MintTx{})
}
