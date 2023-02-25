package main

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"

	"github.com/mehdi124/blockcherry/core"
	"github.com/mehdi124/blockcherry/crypto"
	"github.com/mehdi124/blockcherry/network"
	"github.com/sirupsen/logrus"
)

func main() {

	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			//trRemote.SendMessage(trLocal.Addr(), []byte("Hello world"))
			if err := sendTransaction(trRemote, trLocal.Addr()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	server := network.NewServer(opts)
	server.Start()
}

func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(10000000)), 10))
	tx := core.NewTransaction(data)

	tx.Sign(privKey)

	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())
	return tr.SendMessage(to, msg.Bytes())

}
