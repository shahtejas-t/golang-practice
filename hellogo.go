package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

var pl = fmt.Println

func main() {
	pl("PL TYPEOF : ", reflect.TypeOf(pl))
	pl("Name ?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}
