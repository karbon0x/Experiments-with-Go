package main

import (
	"log"
	"net"
	"net/rpc"
	"os/exec"
)

type Fin struct {
	T string
}

type Arith2 []byte

://Trying to read the file and do GREP

func (t *Arith2) Readi(arg *Fin, reply *[]byte) error {
	cmd := exec.Command("grep", *arg, "test.txt")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	go io.Copy(*reply, stdout)

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
		go rpc.ServeConn(conn)
	}
}
