package main

import (
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	for _, site := range sites {
		wg.Add(1)

		go func(site string) {
			defer wg.Done()

			res, err := http.Get(site)

			if err != nil {
			}

			io.WriteString(os.Stdout, res.Status+"\n")
		}(site)
	}

	wg.Wait()
}
