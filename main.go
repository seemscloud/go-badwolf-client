package main

import (
	"seems.cloud/badwolf/client/cmd/protocol"
)

func main() {
	go protocol.PingHandler()

	select {}
}
