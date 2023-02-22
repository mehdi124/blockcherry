package core

import (
	"crypto/sha256"

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
