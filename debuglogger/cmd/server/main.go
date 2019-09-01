package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var (
	addr = ":19999"
)

func init() {
	flag.StringVar(&addr, "-a", addr, "addr")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Can't accept conn: %s\n", err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	log.Printf("Received %s\n", conn.RemoteAddr().String())

	io.Copy(os.Stdout, conn)
}
