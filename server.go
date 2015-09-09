package main

import (

	"fmt"
	"io/ioutil"
	"net/http"
	"net/rpc"

)


type Fin struct {
	T string
}







type Arith2 []byte


func (t *Arith2) Readi(arg* Fin, reply *string) error{
    data, err := ioutil.ReadFile(arg.T)
      if err!=nil{
          return nil
      }
	  *reply = string(data)

     return nil
}

func main() {

	arith2 := new(Arith2)

	rpc.Register(arith2)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
