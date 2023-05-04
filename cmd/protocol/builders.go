package protocol

import (
	"fmt"
	"os"
)

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Failed to get hostname")
	}
	return name
}

func PingPongBuilder() []byte {
	var message PingPong

	copy(message.Type[:], PingPongType)
	copy(message.Hostname[:], getHostname())

	return pingPongBytes(message)
}

func pingPongBytes(p PingPong) []byte {
	var buf []byte
	buf = append(buf, p.Type[:]...)
	buf = append(buf, p.Hostname[:]...)
	return buf
}
