package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dustin/go-humanize"
)

func save(w http.ResponseWriter, req *http.Request) {
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %v", err)
	}

	fmt.Println(humanize.Bytes(uint64(len(bytes))))

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/logs", save)

	http.ListenAndServe(":8090", nil)
}
