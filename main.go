package main

import (
	"fmt"

	"github.com/ear7h/parsepg/parser"
)

func main() {
	fmt.Println("hello")

	h, err := parser.ParseFile("test_queries.sql")
	if err != nil {
		panic(err)
	}

	queries, ok := h.([]parser.Query)
	if !ok {
		panic("no tokens received")
	}

	fmt.Println(queries)
	for _, q := range queries {
		fmt.Println("q: ", q)
	}
}
