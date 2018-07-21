package main

import (
	"fmt"
	"os"
	"strings"
)

var input *os.File
var output *os.File

var T int
var C []Case

type Case struct {
	A string
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.A = readString()
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...\n")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		res := solve(c.A)
		fmt.Printf("Case #%d: %s\n", i+1, res)
	}
}

func solve(A string) string {
	firstLetterRepeated := -1

	for i := 1; i < len(A); i++ {
		if A[i] == A[0] {
			firstLetterRepeated = i
			break
		}
	}

	if firstLetterRepeated == -1 {
		return "Impossible"
	}

	res := A[:firstLetterRepeated] + A

	if strings.HasPrefix(res, A) {
		return "Impossible"
	}

	return res
}

func readInt() int {
	var i int
	fmt.Fscanf(os.Stdin, "%d", &i)
	return i
}

func readString() string {
	var str string
	fmt.Fscanf(os.Stdin, "%s", &str)
	return str
}
