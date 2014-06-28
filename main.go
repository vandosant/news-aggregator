package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"
)

type Item struct {
    Title string
    URL string
}

type Response struct {
    Data struct {
        Children []struct {
            Data Item
        }
    }
}

func main() {
    response, err := http.Get("http://reddit.com/r/golang.json")
    if err != nil {
	log.Fatal(err)
    }

    if response.StatusCode != http.StatusOK {
	log.Fatal(response.Status)
    }

    resp := new(Response)
    err = json.NewDecoder(response.Body).Decode(resp)

    for _, child := range resp.Data.Children {
        fmt.Println(child.Data.Title)
    }
}