package protocol

type DataTransfer struct {
	Type     [3]byte
	Size     [5]byte
	Checksum [64]byte
	Payload  []byte
}

type PingPong struct {
	Type    [3]byte
	Payload [5]byte
}
