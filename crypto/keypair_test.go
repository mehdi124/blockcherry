package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPairSignVerifySuccess(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	//address := pubKey.Address()
	msg := []byte("Hello world")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	res := sig.Verify(pubKey, msg)
	assert.True(t, res)
}

func TestKeyPairSignVerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()

	msg := []byte("Hello world")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPubKey, msg))
	assert.False(t, sig.Verify(otherPubKey, []byte("fail test")))
}
