package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNum := rand.Intn(maxNum)
	fmt.Println("num is ", secretNum)

	fmt.Println("input guess")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("err", err)
		return
	}
	input = strings.TrimSuffix(input, "\n")
	guess, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("err input", err)
		return
	}
	fmt.Println("your guess", guess)

}
