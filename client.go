package main

import (
    "fmt"
    "log"
    "net/rpc"
    "os"
    "strings"
)

type Fin struct {
	T,T2 string
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
      var A[]byte



arg := Fin{input,input2}
    err = client.Call("Arith2.Readi", arg, &A)
if err != nil {
     log.Fatal("Read error", err)
 }


fmt.Println(string(A))
str := strings.Split(string(A),"\n")
for i := 0; i < len(str); i++ {
    if strings.Contains(str[i], "ftp"){
            fmt.Println(str[i])
        }
    }
}
