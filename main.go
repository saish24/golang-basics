package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		panic(err)
	}

	connChan := make(chan struct{}, 3)

	for {
		log.Println("waiting for client to connect")
		conn, err := listener.Accept()
		connChan <- struct{}{}
		if err != nil {
			panic(err)
		}
		go do(conn, connChan)
	}
}

func do(conn net.Conn, connChan <-chan struct{}) {
	defer func() {
		conn.Close()
		log.Println("request completed")
		<-connChan
	}()
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("processing request")
	time.Sleep(8 * time.Second)

	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World 🌏\r\n"))
	if err != nil {
		log.Fatal(err)
	}
}
