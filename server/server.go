package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	recAddr := ":5001"
	sndAddr := ":5002"

	// Receiver
	tcpRec, err := net.Listen("tcp", recAddr)
	if err != nil {
		panic(err)
	}

	// Sender
	tcpSnd, err := net.Listen("tcp", sndAddr)
	if err != nil {
		panic(err)
	}

	for {
		msg := make(chan string)
		c1, err := tcpRec.Accept()
		if err != nil {
			continue
		} else {
			go recHandler(c1, msg)
		}

		// message := <-msg
		c2, err := tcpSnd.Accept()
		if err != nil {
			continue
		}
		go sndHandler(c2, msg)
	}

}

func recHandler(conn net.Conn, msg chan string) {
	fmt.Println("Receiver connected...")
	_, err := conn.Write([]byte("OK"))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("Wrote to sending server")

	data := make([]byte, 1024)

	_, err = conn.Read(data)
	if err != nil {
		println("Read from server failed:", err.Error())
		os.Exit(1)
	}
	time.Sleep(2 * time.Second)

	msg <- string(data)
	conn.Close()
}

func sndHandler(conn net.Conn, msg chan string) {
	conn.Write([]byte(<-msg))
	conn.Close()
}
