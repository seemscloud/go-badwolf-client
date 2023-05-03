package protocol

import (
	"fmt"
	"net"
	"time"
)

func connCreate() net.Conn {
	for {
		conn, err := net.Dial("tcp", ":11112")
		if err != nil {
			fmt.Println("Error dial connection:", err)
			time.Sleep(time.Second)
		} else {
			return conn
		}
	}
}

func connWrite(conn *net.Conn, data []byte) {
	_, err := (*conn).Write(data)
	if err != nil {
		fmt.Println("Failed sending data:", err)

		if netErr, ok := err.(*net.OpError); ok && netErr.Err.Error() == "write: broken pipe" {
			fmt.Println("Broken pipe: ", err)

			err := (*conn).Close()
			if err != nil {
				fmt.Println("Failed to close connection: ", err)
			}

			*conn = connCreate()
		}

	}
}
