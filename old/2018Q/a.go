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
	N, K, V int
	A       []string
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N, c.K, c.V = readInt(), readInt(), readInt()
		for j := 0; j < c.N; j++ {
			c.A = append(c.A, readString())
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...\n")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		res := solve(c.N, c.K, c.V, c.A)
		fmt.Printf("Case #%d: %s\n", i+1, res)
	}
}

func solve(N, K, V int, A []string) string {
	previousVisits := V - 1
	previousVisits = previousVisits % N

	visited := (K * previousVisits) % N
	notVisited := N - visited

	fmt.Fprintf(os.Stderr, "previousVisits: %d, visited: %d, notVisited: %d\n",
		previousVisits, visited, notVisited)

	res := []string{}

	if K > notVisited {
		seeAgain := K - notVisited
		for i := 0; i < seeAgain; i++ {
			res = append(res, A[i])
		}
	}

	for i := 0; i < K && visited+i < N; i++ {
		res = append(res, A[visited+i])
	}

	return strings.Join(res, " ")
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
