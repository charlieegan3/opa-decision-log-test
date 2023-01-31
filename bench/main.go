package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	input, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000000; i++ {
		resp, err := http.Post(
			"http://localhost:8181/v0/data/example/allow",
			"application/json",
			bytes.NewReader(input),
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(i, resp.StatusCode)
	}
}
