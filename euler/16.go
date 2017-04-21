package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func powerDigitSum() (int, error) {
	n := new(big.Int)
	n.Exp(big.NewInt(2), big.NewInt(1000), nil)
	sum := 0
	for _, c := range n.String() {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, fmt.Errorf("weird number")
		}
		sum += digit
	}
	return sum, nil
}
