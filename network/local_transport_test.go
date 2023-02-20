package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.peers[trb.addr, trb])
	assert.Equal(t, trb.peers[tra.addr, tra])
}

func TestSendMessage(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.peers[trb.addr, trb])
	assert.Equal(t, trb.peers[tra.addr, tra])
	msg := []byte("Hello world")
	assert.Nil(tra.SendMessage(trb.addr, msg))

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rcp.Payload, tra.addr)
}
