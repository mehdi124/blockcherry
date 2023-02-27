package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"

	"github.com/mehdi124/blockcherry/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(header *Header) types.Hash {
	h := sha256.Sum256(header.Bytes())
	return types.Hash(h)
}

type TxHasher struct{}

func (TxHasher) Hash(tx *Transaction) types.Hash {
	return types.Hash(sha256.Sum256(tx.Data))
}

// Hash will hash the whole bytes of the TX no exception.
func (TxHasher) Hash(tx *Transaction) types.Hash {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, tx.Data)
	binary.Write(buf, binary.LittleEndian, tx.To)
	binary.Write(buf, binary.LittleEndian, tx.Value)
	binary.Write(buf, binary.LittleEndian, tx.From)
	binary.Write(buf, binary.LittleEndian, tx.Nonce)

	return types.Hash(sha256.Sum256(buf.Bytes()))
}
