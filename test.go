package main

import(
    "fmt"
)


func call(p int) {
    fmt.Println(p)
    p = 12
    fmt.Println(p)
}

func main()  {
    a:= 23
    var b string  = "Hello"
    fmt.Println(a)
    fmt.Println(b)
    call(a)
    fmt.Println(a)
}
