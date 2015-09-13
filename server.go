package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	//"io/ioutil"
	"os/exec"
)

type Fin struct {
	T, T2 string
}

type Arith2 []byte

func (t *Arith2) Readi(arg *Fin, reply *[]byte) error {
	var err error
	fmt.Println(arg.T2)

	cmd := exec.Command("grep", arg.T2, arg.T)

	*reply, err = cmd.Output()
	//data, err := ioutil.ReadFile(arg.T)
	if err != nil {
		return nil
	}
	//*reply = data
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
