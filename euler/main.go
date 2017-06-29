package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(amicableNumbers(10000))
	// fmt.Println("maximum for triangle : ", maximum)
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}
