package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	filenameSource := "testa.txt"
	filenameDest := "testb.txt"
	serverAddress := "localhost:8080"
	sendFile(filenameSource, serverAddress)
	grabFile(filenameDest, serverAddress)
}

func sendFile(filename, serverAddress string) {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	inFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inFile.Close()

	_, err = io.Copy(conn, inFile)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func grabFile(filename, serverAddress string) {
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = io.Copy(outFile, conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}
