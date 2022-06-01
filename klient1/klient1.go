package main

import (
	"net"
	"os"
	"strings"
)

func main() {

	var hemmelighet = "Møte i Ålesund 1. juni kl. 25:59"

	// fmt.Println(encrypt(hemmelighet))

	servAddr := "localhost:5001"
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

	_, err = conn.Write([]byte(encrypt(hemmelighet)))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
}

func encrypt(text string) string {

	var alfabet = strings.Split("abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ .:1234567890", "")
	var caesar = strings.Split("efghijklmnopqrstuvwxyzæøåabcdEFGHIJKLMNOPQRSTUVWXYZÆØÅABCD .:1234567890", "")

	var result = ""
	var letters = strings.Split(text, "")

	for i := 0; i < len(letters); i++ {
		index := find(alfabet, letters[i])
		result = result + caesar[index]
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
