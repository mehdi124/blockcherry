package main

import (
	"github.com/mehdi124/blockcherry/network"
)

func main() {

	trLocal := network.NewLocalTransport("LOCAL")

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	server := network.NewServer(opts)
	server.Start()
}
