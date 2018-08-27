package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("init")
	t := time.Now()
	fmt.Println(int(t.Month()))
}
