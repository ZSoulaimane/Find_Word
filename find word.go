package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" + "lazy cat jump again and again and agains"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}

			if q == w {
				fmt.Println("#\n", i+1, w)

				break queries
			}

		}
	}
}
