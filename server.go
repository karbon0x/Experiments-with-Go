package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"

	"net/rpc"
	"os/exec"
	//"bytes"
)

type Fin struct {
	T, T2 string
}

type Arith2 []byte

func (t *Arith2) Readi(arg *Fin, reply *[]byte) error {
	var err error
	fmt.Println(arg.T)
	//TODO add file argument for grep function
	cmd := exec.Command("grep", arg.T2, arg.T)
	*reply, err = cmd.Output()
	//	fmt.Println(string(*reply))
	if err != nil {
		return err
	}
	return nil
}

func main() {

	arith2 := new(Arith2)

	rpc.Register(arith2)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1235", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
