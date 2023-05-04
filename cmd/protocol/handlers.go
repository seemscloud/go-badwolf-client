package protocol

import "time"

func PingHandler() {
	conn := createConnection()
	for {
		messageBytes := PingPongBuilder()
		sendData(&conn, messageBytes)

		time.Sleep(time.Second)
	}
}
