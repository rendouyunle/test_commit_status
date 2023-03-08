package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Format(time.RFC3339)
	fmt.Println("execute time: ", now)
}
