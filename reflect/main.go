package main

import (
	"fmt"
	"reflect"
)

func main(){
    var i = 10
    var s = "Viet Nam"

    fmt.Println(reflect.TypeOf(i))
    fmt.Println(reflect.TypeOf(s))
}