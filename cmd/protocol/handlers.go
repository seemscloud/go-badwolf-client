package protocol

func PingHandler() {
	conn := connCreate()
	for {
		connWrite(&conn, PingPongBuilder())
	}
}
