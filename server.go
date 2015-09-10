package main

import (


	"io/ioutil"
	"net"
	"net/rpc"
	"log"

)


type Fin struct {
	T string
}

type Arith2 []byte


func (t *Arith2) Readi(arg* Fin, reply *[]byte) error{
    data, err := ioutil.ReadFile(arg.T)
      if err!=nil{
          return nil
      }
		*reply = data

     return nil
}

func main() {

	arith2 := new(Arith2)
	rpc.Register(arith2)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
	    log.Fatal("listen error:", e)
	}
	    for {
	        conn, err := l.Accept()
			 if err != nil {
			    log.Fatal("listen error:", err)
			}
			rpc.ServeConn(conn)
	    }
}
