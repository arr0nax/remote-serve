package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func get(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got / request at laptop\n")
	io.WriteString(w, "This is my website, served from laptop!\n")
}

func main() {
    http.HandleFunc("/", get)
    err := http.ListenAndServe(":3333", nil)
    fmt.Printf("server started on laptop, port 3333")

    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed on laptop\n")
    } else if err != nil {
        fmt.Printf("error starting server on laptop: %s\n", err)
		os.Exit(1)
    }
}
