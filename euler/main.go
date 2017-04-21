package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(maxPathSum(triangleInput))
	// fmt.Println("maximum for triangle : ", maximum)
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}
