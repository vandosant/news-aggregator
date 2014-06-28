package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    response, err := http.Get("http://reddit.com/r/golang.json")
    if err != nil {
	log.Fatal(err)
    }

    if response.StatusCode != http.StatusOK {
	log.Fatal(response.Status)
    }

    _, err = io.Copy(os.Stdout, response.Body)
    if err != nil {
	log.Fatal(err)
    }
}