package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const location = "flock server"
const port = 3333

func get(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got / request at %v\n", location)
	io.WriteString(w, "This is my website, served from " + location +"!\n")
}

func main() {
    fmt.Printf("server started on %v, port %v\n", location, port)
    http.HandleFunc("/", get)
    err := http.ListenAndServe(":3333", nil)

    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed on %v\n", location)
    } else if err != nil {
        fmt.Printf("error starting server on %v: %s\n", location, err)
		os.Exit(1)
    }
}
