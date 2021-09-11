// OMDb is used to provide movie data
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const (
	omdbURL string = "http://www.omdbapi.com"
)

// Movie is the JSON structure returned by OMDb (only some field are considered)
type Movie struct {
	Response string `json:"Response"`
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Release  string `json:"Release"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Poster   string `json:"Poster"`
	Type     string `json:"Type"`
}

// main is the entry point of the program
func main() {
	// Check parameters
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <apikey> <name>\n", os.Args[0])
		os.Exit(1)
	}

	// Get movie data according to the search tag
	movie, err := getMovieData(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get movie data: %v\n", err)
		os.Exit(2)
	}

	// Check server response
	if movie.Response != "True" {
		fmt.Fprintf(os.Stderr, "Movie not found\n")
		os.Exit(3)
	}

	// Download the poster
	ext := filepath.Ext(movie.Poster)
	name := movie.Title + ext
	err = downloadPoster(movie.Poster, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to download the poster: %v\n", err)
		os.Exit(4)
	}
}

// getMovieData gets data about a movie identified by its title
func getMovieData(apikey, name string) (*Movie, error) {

	// Build the request
	values := make(url.Values)
	values.Add("apikey", apikey)
	values.Add("t", name)
	urlRequest := omdbURL + "/?" + values.Encode()

	// Send the request to the server
	resp, err := http.Get(urlRequest)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("read query failed: %s", resp.Status)
	}

	// Decode the response
	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}

	return &movie, nil
}

// downloadPoster downloads the poster of a movie
func downloadPoster(poster, name string) error {

	// Send the request to the server
	resp, err := http.Get(poster)
	if err != nil {
		return err
	}

	// We must close resp.Body on all execution paths.
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("read query failed: %s", resp.Status)
	}

	// Open the output file
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	// We must close the file on all execution paths
	defer file.Close()

	// Create the file writer
	writer := bufio.NewWriter(file)

	// Copy poster into the file
	_, err = writer.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	// Don't forget to flush
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
