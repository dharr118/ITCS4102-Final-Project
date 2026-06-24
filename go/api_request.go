package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.github.com")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
}
