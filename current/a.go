package main

import (
	"fmt"
	"os"
)

var input *os.File
var output *os.File

var T int
var C []Case

type Case struct {
	N     int
	Line1 []rune
	Line2 []rune
	Line3 []rune
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N = readInt()
		c.Line1 = []rune(readString())
		c.Line2 = []rune(readString())
		c.Line3 = []rune(readString())
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...\n")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		res := solve(c.N, c.Line1, c.Line2, c.Line3)
		fmt.Printf("Case #%d: %d\n", i+1, res)
	}
}

const modulo = 1000000007

func solve(N int, Line1, Line2, Line3 []rune) int {
	if N == 0 || N%2 != 0 {
		return 0
	}

	// check first and last rows
	if Line1[0] == '#' || Line2[0] == '#' {
		return 0
	}
	if Line2[N-1] == '#' || Line3[N-1] == '#' {
		return 0
	}

	total := 1

	for i := 1; i < N-1; i += 2 {
		// impossible if the middle line is blocked somewhere
		if Line2[i] == '#' || Line2[i+1] == '#' {
			return 0
		}

		possiblities := 2

		if Line1[i] == '#' || Line1[i+1] == '#' {
			possiblities--
		}
		if Line3[i] == '#' || Line3[i+1] == '#' {
			possiblities--
		}

		if possiblities == 0 {
			return 0
		}

		total = total * possiblities

		if total > modulo {
			total = total % modulo
		}
	}

	return total
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
