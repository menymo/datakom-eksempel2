package main

import (
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	servAddr := "localhost:5002"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	msg := make([]byte, 1024)

	_, err = conn.Read(msg)
	if err != nil {
		println("Read from server failed:", err.Error())
		os.Exit(1)
	}
	time.Sleep(2 * time.Second)

	decrypt(string(msg))
}

func decrypt(text string) string {

	var alfabet = strings.Split("abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ .:1234567890", "")
	var caesar = strings.Split("efghijklmnopqrstuvwxyzæøåabcdEFGHIJKLMNOPQRSTUVWXYZÆØÅABCD .:1234567890", "")

	var result = ""
	var letters = strings.Split(text, "")

	for i := 0; i < 32; i++ {
		index := find(caesar, letters[i])
		result = result + alfabet[index]
	}

	return result

}

func find(a []string, s string) int {
	for i, n := range a {
		if s == n {
			return i
		}
	}
	return len(a)
}
