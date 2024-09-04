package main

import "fmt"

func main() {
	i := 13
	sum := 0
	k := 0

	for k < i {
		k = k + 1
		sum = sum + k
	}

	fmt.Println(sum)
}
