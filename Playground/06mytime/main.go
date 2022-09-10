package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hey! Tiem now is")
	t := time.Now()
	fmt.Println(t.Format("01-09-2006 Monday"))

}
