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
	serverAddress := "127.0.0.1:8080"

	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}
	defer listener.Close()

	receiveFile(listener, filenameDest)
	sendFile(listener, filenameSource)
}

func receiveFile(listener net.Listener, dest string) {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		return
	}
	defer conn.Close()

	outFile, err := os.Create(dest)
	if err != nil {
		fmt.Println("Error creating file: ", err.Error())
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, conn)
	if err != nil {
		fmt.Println("Error copying file: ", err.Error())
	}
}

func sendFile(listener net.Listener, source string) {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		return
	}
	defer conn.Close()

	inFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Error opening file: ", err.Error())
		return
	}
	defer inFile.Close()

	_, err = io.Copy(conn, inFile)
	if err != nil {
		fmt.Println("Error copying file: ", err.Error())
		return
	}
}
