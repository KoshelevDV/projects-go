package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func isprime(n int) bool {
	switch {
	case n == 2:
		return true
	case n == 3:
		return true
	case n%2 == 0:
		return false
	case n%3 == 0:
		return false
	}

	i := 5
	w := 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}

	return true
}

func sliceisprime(n []string) (res []string) {
	for _, elem := range n {
		temp, err := strconv.Atoi(elem)
		if err != nil {
			log.Printf("%s is not a number. %v", elem, err)
		} else {
			if isprime(temp) {
				res = append(res, elem)
			}
		}
	}
	return
}

func main() {

	args := os.Args[1:]
	fmt.Println(sliceisprime(args))
	fmt.Println(isprime(11))
}
