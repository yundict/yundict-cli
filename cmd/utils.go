package cmd

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Download file from url
func downloadFile(url string, outPath string) {

	// append extension if not exists(check outPath ends with "/")
	if strings.HasSuffix(outPath, "/") {
		r, _ := http.NewRequest("GET", url, nil)
		filename := filepath.Base(r.URL.Path)
		outPath = filepath.Join(outPath, filename)
	}

	// Create the file
	out, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
