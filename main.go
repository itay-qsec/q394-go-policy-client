// A Go HTTPS client that pins its TLS policy to classical X25519 only.
// Acceptance-test fixture for Q-394.
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion:       tls.VersionTLS13,
			CurvePreferences: []tls.CurveID{tls.X25519},
		},
	}
	client := &http.Client{Transport: transport}
	resp, err := client.Get("https://api.github.com/zen")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
