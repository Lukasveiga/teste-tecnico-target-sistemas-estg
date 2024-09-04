package main

import (
	"errors"
	"fmt"
	"log"
)

func fibonacci(num int) (bool, error) {
	second := 1
	third := 1
	fibNumber := 0

	if num < 0 {
		return false, errors.New("input cannot be negative")
	}

	if num == fibNumber {
		return true, nil
	}

	for i := 0; i < num-2; i++ {
		fibNumber = second + third
		second = third
		third = fibNumber

		if fibNumber == num {
			return true, nil
		}
	}

	return false, nil
}

func input() (int, error) {
	var n int
	fmt.Print("Input: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return -1, err
	}

	return n, nil
}

func main() {
	n, err := input()

	if err != nil {
		log.Fatal(err)
	}

	contains, err := fibonacci(n)

	if err != nil {
		log.Fatal(err)
	}

	if contains {
		fmt.Printf("%d pertence a sequência de fibonacci\n", n)
	} else {
		fmt.Printf("%d não pertence a sequência de fibonacci\n", n)
	}
}
