package protocol

func PingHandler() {
	conn := connCreate()
	for {
		messageBytes := PingPongBuilder()
		connWrite(&conn, messageBytes)
	}
}
