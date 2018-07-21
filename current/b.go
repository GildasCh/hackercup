package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input *os.File
var output *os.File

var T int
var C []Case

type Case struct {
	N, K int
	A    []int
	B    []int
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N, c.K = readInt(), readInt()
		for j := 0; j < c.N; j++ {
			c.A = append(c.A, readInt())
			c.B = append(c.B, readInt())
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...\n")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		res := solve(c.N, c.K, c.A, c.B)
		fmt.Printf("Case #%d: %s\n", i+1, res)
	}
}

func solve(N, K int, A, B []int) string {
	pre := preOrder(0, N, A, B)
	post := postOrder(0, N, A, B)

	// fmt.Fprintf(os.Stderr, "pre: %v\n", pre)
	// fmt.Fprintf(os.Stderr, "post: %v\n", post)

	preToPost := make(map[int]int)
	for i := range pre {
		preToPost[pre[i]] = post[i]
	}

	// fmt.Fprintf(os.Stderr, "preToPost: %v\n", preToPost)

	labels := make(map[int]int)

	for k := 1; k <= K; k++ {
		unset := -1

		// find next avalaible label
		for _, i := range pre {
			if _, ok := labels[i]; !ok {
				unset = i
				break
			}
		}

		if unset == -1 {
			return "Impossible"
		}

		// fmt.Fprintf(os.Stderr, "Found unset: %d\n", unset)
		// fmt.Fprintf(os.Stderr, "Setting label %d with %d\n", unset, k)
		labels[unset] = k

		// set all labels that need to be equal
		next := preToPost[unset]
		for {
			if _, ok := labels[next]; ok {
				break
			}
			// fmt.Fprintf(os.Stderr, "Setting label %d with %d\n", next, k)
			labels[next] = k

			next = preToPost[next]
		}
	}

	res := ""

	for i := 1; i <= N; i++ {
		l := labels[i]
		if l == 0 {
			l = 1
		}
		res += strconv.Itoa(l) + " "
	}

	res = strings.TrimSuffix(res, " ")

	return res
}

func preOrder(i, N int, A, B []int) []int {
	parc := []int{i + 1}

	if A[i] != 0 {
		parc = append(parc, preOrder(A[i]-1, N, A, B)...)
	}

	if B[i] != 0 {
		parc = append(parc, preOrder(B[i]-1, N, A, B)...)
	}

	return parc
}

func postOrder(i, N int, A, B []int) []int {
	parc := []int{}

	if A[i] != 0 {
		parc = append(parc, postOrder(A[i]-1, N, A, B)...)
	}

	if B[i] != 0 {
		parc = append(parc, postOrder(B[i]-1, N, A, B)...)
	}

	parc = append(parc, i+1)

	return parc
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
