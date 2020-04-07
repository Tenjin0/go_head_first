package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Hello world")
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(reflect.TypeOf(year))

}
