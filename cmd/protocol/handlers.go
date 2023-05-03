package protocol

import "time"

func PingHandler() {
	conn := connCreate()
	for {
		messageBytes := PingPongBuilder()
		connWrite(&conn, messageBytes)

		time.Sleep(time.Second)
	}
}
