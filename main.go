package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/3stadtCode/001-introduction/client"
	"github.com/3stadtCode/001-introduction/server"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Sending: %q\n", strings.Join(os.Args[1:], " "))
		client.Send(strings.Join(os.Args[1:], " "), fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
		return
	}

	server.Serve(CONN_HOST, CONN_PORT)
}
