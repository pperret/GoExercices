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
	xkcdURL    string = "http://www.xkcd.com"
	xkcdJSON   string = "info.0.json"
	indexFile  string = "xkcd.idx"
	comicsFile string = "xkcd.db"
)

// Number is the main ID of a comic
type Number int

// Numbers is a list of comics (identified by their number)
type Numbers []Number

// WordIndexes is the word indexes (the content of the index file)
type WordIndexes map[string]Numbers

// Comic is the JSON structure returned by XKSD
type Comic struct {
	Number     Number `json:"num"`
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

// Comics is the map of comics (JSON structures returned by XKCD) identified by its number
type Comics map[Number]Comic

// main is the entry point of the program
func main() {
	// Check the parameters count
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <tag> [tags...]\n", os.Args[0])
		os.Exit(1)
	}

	// Load the comics (or downloads them if they are not already stored locally)
	comics, err := loadComics()
	if err != nil {
		fmt.Printf("Unable to get comics: %v\n", err)
		os.Exit(1)
	}

	// Load the index file (or builds it if it does not already exist)
	indexes, err := loadIndexFile(comics)
	if err != nil {
		fmt.Printf("Unable to build indexes: %v\n", err)
		os.Exit(1)
	}

	// Look for entries matching the first tag (using the indexes)
	numbers, ok := (*indexes)[os.Args[1]]
	if ok {
		// Check if the others tags match the entries (using the indexes)
		for _, tag := range os.Args[2:] {
			numbers2, ok := (*indexes)[tag]
			if !ok {
				numbers = Numbers{}
				break
			}
			numbers3 := Numbers{}
			for _, n1 := range numbers {
				for _, n2 := range numbers2 {
					if n1 == n2 {
						numbers3 = append(numbers3, n1)
						break
					}
				}
			}
			numbers = numbers3
		}

		// Display the entries matching all tags
		for _, number := range numbers {
			comic := comics[number]
			fmt.Printf("%s\n", buildURL(comic.Number))
			fmt.Printf("\t%s\n", comic.Transcript)
		}
	}

}

// loadComics load comics or downloads them if they are not already stored locally
func loadComics() (Comics, error) {
	// Reads comics from the local store
	comics, err := readComics()
	if err != nil {
		return nil, err
	}

	// If the comics are not yet downloaded
	if comics == nil {
		// Download comics from the server
		comics, err = downloadComics()
		if err != nil {
			return nil, err
		}

		// Store the comics into the local cache
		err = writeComics(comics)
		if err != nil {
			return nil, err
		}
	}
	return comics, nil
}

// readComics reads comics from the local cache
// A nil value (not an error) is returned when the local cache does not exist
func readComics() (Comics, error) {
	// Open the input file
	file, err := os.Open(comicsFile)
	if err != nil {
		// If the file does not exist, a nil value is returned to indicate that the comics are not yet downloaded
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	// Create the file reader
	reader := bufio.NewReader(file)

	// Decode the file content using the JSON decoder
	var comics Comics
	if err := json.NewDecoder(reader).Decode(&comics); err != nil {
		file.Close()
		return nil, err
	}

	file.Close()
	return comics, nil
}

// writeComics writes comics into the local cache
func writeComics(comics Comics) error {
	// Open the output file
	file, err := os.OpenFile(comicsFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	// Create the file writer
	writer := bufio.NewWriter(file)

	// Encode the file using the JSON encoder
	if err := json.NewEncoder(writer).Encode(comics); err != nil {
		file.Close()
		return err
	}

	// Don't forget to flush
	writer.Flush()

	file.Close()
	return nil

}

// downloadComics downloads all the comics from the server
func downloadComics() (Comics, error) {

	// Get the last entry to determine the max number of the comics
	lastEntry, err := downloadEntry(xkcdURL + "/" + xkcdJSON)
	if err != nil {
		return nil, err
	}

	// Read the comics list
	comics := make(Comics)
	for i := Number(1); i <= lastEntry.Number; i++ {
		// Download one comic (if the required entry does not exist, no error is returned but a nil pointer)
		entry, err := downloadComic(i)
		if err != nil {
			return nil, err
		}

		// Skip non-existing entries
		if entry != nil {
			comics[entry.Number] = *entry
		}
	}

	return comics, nil
}

// downloadComic downloads one comic entry from the server using its number
// An nil value (not an error) is returned if the requested comic does not exist
func downloadComic(number Number) (*Comic, error) {
	return downloadEntry(buildURL(number))
}

// buildURL builds the comic's URL from its number
func buildURL(number Number) string {
	return xkcdURL + "/" + strconv.Itoa(int(number)) + "/" + xkcdJSON
}

// downloadEntry downloads one comic entry from the server using its URL
// An nil value is returned if the requested comic does not exist
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
			return nil, nil
		}
		return nil, fmt.Errorf("read query failed: %s", resp.Status)
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

// loadIndexFile loads index file or creates it if it does not already exist
func loadIndexFile(comics Comics) (*WordIndexes, error) {

	// Read the index file (a nil value is returned if the indexes are not yet built)
	indexes, err := readIndexFile()
	if err != nil {
		return nil, err
	}

	// A nil value is returned if indexes are not yet built
	if indexes == nil {
		// Build the indexes
		indexes, err = buildIndex(comics)
		if err != nil {
			return nil, err
		}

		// Store the indexes in a local file
		err = writeIndexFile(indexes)
		if err != nil {
			return nil, err
		}
	}
	return indexes, nil
}

// readIndexFile reads indexes from the local cache file
// A nil value (not an error) is returned if there is cache file
func readIndexFile() (*WordIndexes, error) {

	// Open the input file
	file, err := os.Open(indexFile)
	if err != nil {
		// If the file does not exist, returns a nil value to indicate that the indexes are not yet built
		if os.IsNotExist(err) {
			return nil, nil
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

// buildIndex creates the indexes from comics
func buildIndex(comics Comics) (*WordIndexes, error) {
	// Builds the map
	indexes := make(WordIndexes)

	// Loop on comic list
	for _, comic := range comics {

		// Get words in the comic transcript
		words, err := ScanWords(comic.Transcript)
		if err != nil {
			return nil, err
		}

		// Loop on comic words
		for word := range words {
			// Appends each word to the map
			indexes[word] = append(indexes[word], comic.Number)
		}
	}
	return &indexes, nil
}

// ScanWords scans words in a string
func ScanWords(text string) (map[string]bool, error) {

	// Create the scanner
	scanner := bufio.NewScanner(strings.NewReader(text))

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Loop on words and store them into a map to manage unicity
	words := make(map[string]bool)
	for scanner.Scan() {
		word := scanner.Text()
		words[word] = true
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// WriteIndexFile writes indexes into the local file
func writeIndexFile(indexes *WordIndexes) error {

	// Open the output file
	file, err := os.OpenFile(indexFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	// Create the file writer
	writer := bufio.NewWriter(file)

	// Encode the file using the JSON encoder
	if err := json.NewEncoder(writer).Encode(indexes); err != nil {
		file.Close()
		return err
	}

	// Don't forget to flush
	writer.Flush()

	file.Close()
	return nil
}
