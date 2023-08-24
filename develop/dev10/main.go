package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go run . [--timeout=<timeout>] <host> <port>")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	err := telnetClient(host, port, *timeout)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func telnetClient(host string, port string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	go func() {
		io.Copy(os.Stdout, conn)
	}()

	io.Copy(conn, os.Stdin)

	return nil
}
