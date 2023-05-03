package protocol

func PingPongBuilder() []byte {
	var message PingPong

	copy(message.Type[:], PingPongType)
	copy(message.Payload[:], PingMessage)

	return pingPongBytes(message)
}

func pingPongBytes(p PingPong) []byte {
	var buf []byte
	buf = append(buf, p.Type[:]...)
	buf = append(buf, p.Payload[:]...)
	return buf
}
