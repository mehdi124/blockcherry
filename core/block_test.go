package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/mehdi124/blockcherry/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {

	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     983245,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	//	hDecode := &Header{}
	assert.Nil(t, h.DecodeBinary(buf))
	//	assert.Equal(t, h, hDecode)

}

func TestBlock_Encode_Decode(t *testing.T) {

	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     983245,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)

}

func TestBlockHash(t *testing.T) {

	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     983245,
		},
		Transactions: []Transaction{},
	}

	h := b.Hash()
	fmt.Println("block hash", h)
	assert.False(t, h.IsZero())

}
