package main

import (
    "fmt"
    "log"
    "net/rpc"
    "os"
    "strings"
)

type Fin struct {
	T string
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "server")
        os.Exit(1)
    }

    serverAddress := os.Args[1]
    client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

var A[]byte
arg := Fin{"client.go"}
    err = client.Call("Arith2.Readi", arg, &A)
if err != nil {
     log.Fatal("Read error", err)
 }
str := strings.Split(string(A),"\n")
for i := 0; i < len(str); i++ {
    if strings.Contains(str[i], "le"){
        fmt.Println(str[i])
    }
}
}
