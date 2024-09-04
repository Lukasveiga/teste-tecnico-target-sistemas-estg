package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func input() (string, error) {
	var s string
	fmt.Print("Input: ")
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return "", err
	}

	if len(s) == 0 {
		return "", errors.New("Input cannot be empty")
	}

	return s, nil
}

func reverse(s string) string {
	r := ""
	words := strings.Split(s, "")

	for i := len(words) - 1; i >= 0; i-- {
		r += words[i]
	}

	return r
}

func main() {
	s, err := input()

	if err != nil {
		log.Fatal(err)
	}

	r := reverse(s)
	fmt.Println(r)
}
