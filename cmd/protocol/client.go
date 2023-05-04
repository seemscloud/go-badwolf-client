package protocol

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

const (
	serverAddr  = ":11112"
	retryPeriod = 1 * time.Second
)

func createConnection() net.Conn {
	for {
		tlsConfig := tls.Config{InsecureSkipVerify: true}
		conn, err := tls.Dial("tcp", serverAddr, &tlsConfig)
		if err != nil {
			fmt.Printf("Failed to connect to server: %v\n", err)
			time.Sleep(retryPeriod)
			continue
		}
		return conn
	}
}

func sendData(conn *net.Conn, data []byte) {
	_, err := (*conn).Write(data)
	if err != nil {
		fmt.Printf("Failed to send data: %v\n", err)
		if isBrokenPipeError(err) {
			fmt.Println("Broken pipe error encountered, reconnecting to server...")
			_ = (*conn).Close()
			*conn = createConnection()
			return
		}
	}
}

func isBrokenPipeError(err error) bool {
	opErr, ok := err.(*net.OpError)
	if !ok {
		return false
	}

	sysErr, ok := opErr.Err.(*net.OpError)
	if !ok {
		return false
	}

	return sysErr.Err.Error() == "write: broken pipe"
}
