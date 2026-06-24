package main

import (
    "fmt"
    "net/http"
)

func main() {
    response, err := http.Get("https://api.github.com")
    if err != nil {
        fmt.Println("Request error:", err)
        return
    }
    defer response.Body.Close()

    fmt.Println("Status code:", response.StatusCode)
    fmt.Println("Content type:", response.Header.Get("Content-Type"))
}
