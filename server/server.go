package server

import (
	"fmt"
	"net"
	"os"
)

type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
}

func Serve(host, port string) error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return err
	}
	defer l.Close()
	mainLoop(l)
	return nil
}

func mainLoop(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn Conn) error {
	defer conn.Close()

	// Everything above the buffer size is discarded
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return err
	}

	fmt.Printf("Received: %q\n", string(buf))

	conn.Write([]byte("Message received."))
	return nil
}
