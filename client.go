package main

import (
    "fmt"
    "log"
    "net/rpc"
    "os"

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

for 1>0 {
var A string
arg := Fin{"client.go"}
    err = client.Call("Arith2.Readi", arg, &A)
if err != nil {
     log.Fatal("Read error", err)
 }
 fmt.Println(A)

}

}
