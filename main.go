//go:generate pigeon -o pg_gen.go pg.peg

package main

import (
	"fmt"
)

type Query string

type Queries []Query

func main() {
	fmt.Println("hello")

	h, err := ParseFile("test_queries.sql")
	if err != nil {
		panic(err)
	}

	queries, ok := h.([]Query)
	if !ok {
		panic("no tokens received")
	}

	fmt.Println(queries)
	for _, q := range queries {
		fmt.Println("q: ", q)
	}
}
