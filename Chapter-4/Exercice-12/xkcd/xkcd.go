// xkcd is used to seach in xkcd comics
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	xkcdURL   string = "http://www.xkcd.com"
	xkcdJSON  string = "info.0.json"
	fileCache string = "xkcd.db"
)

// Comic is the JSON structure returned by XKSD
type Comic struct {
	Number     int    `json:"num"`
	Title      string // `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Year       string // `json:"year"`
	Month      string // `json:"month"`
	Day        string // `json:"day"`
	Link       string // `json:"link"`
	News       string // `json:"news"`
	Transcript string // `json:"transcript"`
	Alt        string // `json:"alt"`
	Img        string // `json:"img"`
}

// Comics is the list of comics (JSON structures returned by XKCD)
type Comics []Comic

// Integers is the list of comics (identified by their number)
type Integers []int

// WordIndexes is the word indexes (the content of the index file)
type WordIndexes map[string]Integers

// main is the entry point of the program
func main() {
	// Check parameters
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <tag>\n", os.Args[0])
		os.Exit(1)
	}

	// Load the index file (or create it if it does not already exist)
	indexes, err := loadIndexFile()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Look for a matching entry in the index
	numbers, ok := (*indexes)[os.Args[1]]
	if ok == true {
		// Read the comics from XKSD
		for _, number := range numbers {
			comic, err := downloadComic(number)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", buildURL(comic.Number))
			fmt.Printf("\t%s\n", comic.Transcript)
		}
	}
}

// loadIndexFile loads index file or creates it if it does not already exist
func loadIndexFile() (*WordIndexes, error) {

	// Read index file (an empty list is returned if the entries are not already downloaded)
	indexes, err := readIndexFile()
	if err != nil {
		return nil, err
	}

	// If the index was not already loaded, comics are downloaded from the server
	if len(*indexes) == 0 {

		// Download comics from the server
		entries, err := downloadComics()
		if err != nil {
			return nil, err
		}

		// Create the index file from comics
		indexes, err = writeIndexFile(entries)
		if err != nil {
			return nil, err
		}
	}

	return indexes, nil
}

// readIndexFile reads indexes from the local cache file
func readIndexFile() (*WordIndexes, error) {

	// Open the input file
	file, err := os.Open(fileCache)
	if err != nil {
		// If the file does not exist, returns an empty to indicate that the entries are not already downloaded
		if os.IsNotExist(err) {
			return &WordIndexes{}, nil
		}
		return nil, err
	}

	// Create the file reader
	reader := bufio.NewReader(file)

	// Decode the file content using the JSON decoder
	var indexes WordIndexes
	if err := json.NewDecoder(reader).Decode(&indexes); err != nil {
		file.Close()
		return nil, err
	}

	file.Close()
	return &indexes, nil
}

// WriteIndexFile writes entries into the local cache file
func writeIndexFile(comics *Comics) (*WordIndexes, error) {

	// Build the map
	indexes := make(WordIndexes)

	// Loop on comic list
	for _, comic := range *comics {

		// Get words in the comic transcript
		words, err := ScanWords(comic.Transcript)
		if err != nil {
			return nil, err
		}

		// Loop on comic words
		for _, word := range *words {
			// Append each word to the map
			indexes[word] = append(indexes[word], comic.Number)
		}
	}

	// Open the output file
	file, err := os.OpenFile(fileCache, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}

	// Create the file writer
	writer := bufio.NewWriter(file)

	// Encode the file
	if err := json.NewEncoder(writer).Encode(indexes); err != nil {
		file.Close()
		return nil, err
	}

	// Don't forget to flush
	writer.Flush()

	file.Close()
	return &indexes, nil
}

// ScanWords scans words in a string
func ScanWords(text string) (*[]string, error) {

	// Create the scanner
	scanner := bufio.NewScanner(strings.NewReader(text))

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Loop on words
	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		// The word is appended to the list if it is not already present
		found := false
		for _, w := range words {
			if w == word {
				found = true
				break
			}
		}
		if found == false {
			words = append(words, word)
		}
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &words, nil
}

// downloadComics downloads all the comic entries from the server
func downloadComics() (*Comics, error) {

	// Get the last entry to determine the max number of the comics
	lastEntry, err := downloadEntry(xkcdURL + "/" + xkcdJSON)
	if err != nil {
		return nil, err
	}

	// Read the entries list
	var comics Comics
	for i := 1; i <= lastEntry.Number; i++ {
		// Download one comic entry (if the entry does not exist, its number will be 0)
		entry, err := downloadComic(i)
		if err != nil {
			return nil, err
		}

		// Skip non-existing entries (number==0)
		if entry.Number != 0 {
			comics = append(comics, *entry)
		}
	}

	return &comics, nil
}

// downloadComic downloads one comic entry from the server using its number
// An empty value is returned if the requested comic does not exist
func downloadComic(number int) (*Comic, error) {
	return downloadEntry(buildURL(number))
}

// Build the comic's URL from its number
func buildURL(number int) string {
	return xkcdURL + "/" + strconv.Itoa(number) + "/" + xkcdJSON
}

// downloadEntry downloads one comic entry from the server using its URL
// An empty value is returned if the requested comic does not exist
func downloadEntry(url string) (*Comic, error) {

	// Send the request to the server
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()

		// Return an empty entry if the comic is not available
		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("%s is missing\n", url)
			return &Comic{}, nil
		}
		return nil, fmt.Errorf("Read query failed: %s", resp.Status)
	}

	// Decode the response
	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
