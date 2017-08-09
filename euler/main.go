package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	namesScores()
	// fmt.Println("maximum for triangle : ", maximum)
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}
