package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

type myData struct {
	a int
	b float32
	c string
}

func main() {
	dataChannel := make(chan myData)

	go func() {
		mydata := myData{1, 1.1, "abc"}
		dataChannel <- mydata
	}()

	getVal := <-dataChannel
	pl(reflect.TypeOf(getVal))
	pl(getVal.a, "\t", getVal.b, "\t", getVal.c)
}
