package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/mehdi124/blockcherry/crypto"
	"github.com/mehdi124/blockcherry/types"
	"github.com/stretchr/testify/assert"
)

func TestHashBlock(t *testing.T) {
	b := randomBlock(0, types.Hash{})
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	other := crypto.GeneratePrivateKey()
	b.Validator = other.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {

	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height, prevBlockHash)

	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)

	assert.Nil(t, b.Sign(privKey))
	return b
}
