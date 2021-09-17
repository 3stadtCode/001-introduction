package client

import (
	"bufio"
	"fmt"
	"net"
)

func Send(msg string, uri string) error {
	conn, err := net.Dial("tcp", uri)
	if err != nil {
		return err
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	conn.Write([]byte(msg))
	for {
		buf := make([]byte, 1024)
		num, err := reader.Read(buf)

		if err != nil {
			return err
		}

		mensagem := make([]byte, num)
		copy(mensagem, buf)

		fmt.Println(string(mensagem))
	}
}
