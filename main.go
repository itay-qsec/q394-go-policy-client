// A plain Go HTTPS client. Acceptance-test fixture for Q-394 (selector precondition branch).
package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/zen")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
