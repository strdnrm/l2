package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <url>")
		return
	}

	url := os.Args[1]
	err := downloadWebsite(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func downloadWebsite(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download website. Status code: %d", resp.StatusCode)
	}

	filename := getFilename(url)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Website downloaded successfully.")
	return nil
}

func getFilename(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
