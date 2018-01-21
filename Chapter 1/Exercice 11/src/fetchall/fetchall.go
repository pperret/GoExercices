// This program gets main page of the top million web sites
package main

import (
	"archive/zip"
	"bufio"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const urlAlexa string = "http://s3.amazonaws.com/alexa-static/top-1m.csv.zip"
const zipFileName = "top-1m.csv.zip"
const csvFileName string = "top-1m.csv"
const maxFetchs int = 100

func main() {
	start := time.Now()

	// Create the list
	lst := list.New()

	// Get the list (ZIP file)
	res := fetchlist(urlAlexa, zipFileName)
	if res != true {
		os.Exit(1)
	}
	defer os.Remove(zipFileName)

	// Extract CSV list
	res2 := extractlist(zipFileName, csvFileName)
	if res2 != true {
		os.Exit(1)
	}
	defer os.Remove(csvFileName)

	// Parse the CSV file
	res3 := parselist(csvFileName, lst)
	if res3 != true {
		os.Exit(1)
	}

	// Create the channel
	ch := make(chan string)

	// Start URL fetching
	i := 0
	for e := lst.Front(); e != nil; e = e.Next() {
		url := e.Value.(string)
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}
		go fetch(url, ch) // start a goroutine
		i++
		if i >= maxFetchs {
			break
		}
	}

	// Get routine result
	for j := 0; j < i; j++ {
		fmt.Println(<-ch)
		// receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Get the top million web sites
func fetchlist(url string, zipfilename string) bool {

	// Access the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Unable to access 1M list '%s': %v\n", url, err)
		return false
	}
	defer resp.Body.Close()

	// Create the output file
	zipfile, err := os.OpenFile(zipfilename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("Unable to create '%s': %v", zipfilename, err)
		return false
	}
	defer zipfile.Close()

	// Get ZIP file content
	_, err2 := io.Copy(zipfile, resp.Body)
	if err2 != nil {
		fmt.Printf("Unable to store 1M list: %v\n", err2)
		return false
	}

	return true
}

// Extract CSV from ZIP
func extractlist(zipFileName string, csvFileName string) bool {

	// Open ZIP file
	reader, err := zip.OpenReader(zipFileName)
	if err != nil {
		fmt.Printf("Unable to open ZIP file '%s': %v\n", zipFileName, err)
		return false
	}
	defer reader.Close()

	// Check that ZIP file contains only one file
	if len(reader.File) != 1 {
		fmt.Printf("ZIP file contains several files: %d\n", len(reader.File))
		return false
	}

	// Get the unique ZIP subfile
	f := reader.File[0]

	// Open the ZIP subfile
	csvInputFile, err := f.Open()
	if err != nil {
		fmt.Printf("Unable to open ZIP subfile '%s': %v", f.Name, err)
		return false
	}
	defer csvInputFile.Close()

	// Create the output CSV file
	csvOutputFile, err := os.OpenFile(csvFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("Unable to create CSV file '%s': %v", csvFileName, err)
		return false
	}
	defer csvOutputFile.Close()

	// Copy the CSV file
	// Unable to use 'io.Copy' because an unexpected EOF error occurs without any reason
	_, err = io.CopyN(csvOutputFile, csvInputFile, int64(f.UncompressedSize64))
	if err != nil {
		fmt.Printf("Unable to copy CSV file: %v", err)
		return false
	}

	return true
}

// Parse the site list (CSV file)
func parselist(csvFileName string, lst *list.List) bool {

	// Open the CSV file
	f, err := os.Open(csvFileName)
	if err != nil {
		fmt.Printf("Unable to open CSV file %q: %v", csvFileName, err)
		return false
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			lst.PushBack(parts[1])
		} else {
			fmt.Printf("Invalid CSV ligne: %q", line)
		}
	}
	return true
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	// Access to the URL
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error when getting URL: %q, %v", url, err) // send to channel ch
		return
	}
	defer resp.Body.Close()

	// Get URL body (content is ignored)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading: %q, %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %q", secs, nbytes, url)
}
