package protocol

type DataTransfer struct {
	Type     [8]byte
	Size     [5]byte
	Checksum [64]byte
	Payload  []byte
}

type PingPong struct {
	Type     [8]byte
	Hostname [256]byte
}
