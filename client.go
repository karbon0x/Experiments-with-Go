package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
)

type Fin struct {
	T, T2 string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddress := os.Args[1]
	client, err := rpc.Dial("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Print(input)
	var input2 string
	fmt.Scanln(&input2)
	fmt.Print(input2)
	var A []byte

	arg := Fin{input, input2}
	err = client.Call("Arith2.Readi", arg, &A)
	if err != nil {
		log.Fatal("Read error", err)
	}
	i
}
