package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// url -> PLUGIN_URL
	// body -> PLUGIN_BODY
	// method -> PLUGIN_METHOD
	var url = os.Getenv("PLUGIN_URL")
	var method = os.Getenv("PLUGIN_METHOD")
	var bodyStr = os.Getenv("PLUGIN_BODY")

	var client = &http.Client{}

	var body io.Reader
	if method != "GET" && bodyStr != "" {
		body = strings.NewReader(bodyStr)
	}

	var req, err = http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("code: %d", resp.StatusCode)
}
