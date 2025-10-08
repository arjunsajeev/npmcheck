package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

const registryURL = "https://registry.npmjs.org/"

var client = &http.Client{Timeout: 10 * time.Second}

type result struct {
	name      string
	available bool
	err       error
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <package-name> ...\n", os.Args[0])
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for _, name := range os.Args[1:] {
		wg.Go(func() {
			print(check(client, registryURL, name))
		})
	}
	wg.Wait()
}

func check(client *http.Client, url, name string) result {
	resp, err := client.Get(url + name)
	if err != nil {
		return result{name, false, err}
	}
	defer resp.Body.Close()

	return result{name, resp.StatusCode == http.StatusNotFound, nil}
}

func print(r result) {
	if r.err != nil {
		fmt.Printf("⚠️  %s: %v\n", r.name, r.err)
		return
	}
	if r.available {
		fmt.Printf("✅ %s\n", r.name)
	} else {
		fmt.Printf("❌ %s\n", r.name)
	}
}
