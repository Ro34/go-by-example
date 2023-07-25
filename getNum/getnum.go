package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 100
	secretNum := rand.Intn(maxNum)
	fmt.Println("num is ", secretNum)

}
