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
	secretNumber := rand.Intn(maxNum)
	//fmt.Println(secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error! Please try again.")
			continue
		}
		input = strings.TrimSuffix(input, "\r\n")
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}

		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again.")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again.")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}