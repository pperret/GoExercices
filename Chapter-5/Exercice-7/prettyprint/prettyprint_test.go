package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

// URL is the HTML page to test
const URL = "http://golang.org"

// TestPrettyPrint1 tests the prettyPrint function using a memory buffer
func TestPrettyPrint1(t *testing.T) {
	var buffer bytes.Buffer

	// Execute the test
	if err := prettyPrint(&buffer, URL); err != nil {
		t.Fatalf("Error while formatting %v", err)
	}

	// Parse the result
	_, err := html.Parse(&buffer)
	if err != nil {
		t.Fatalf("Error while parsing result: %v", err)
	}
}

// TestPrettyPrint2 tests the prettyPrint function using a temporary file
func TestPrettyPrint2(t *testing.T) {

	// Create the temporary file
	tempfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("Unable to open temp file: %v", err)
	}

	// Execute the test
	if err := prettyPrint(tempfile, URL); err != nil {
		t.Fatalf("Error while formatting %v", err)
	}

	// Parse the result
	_, err = html.Parse(tempfile)
	if err != nil {
		t.Fatalf("Error while parsing result: %v", err)
	}

	// Close the temporary file
	tempfile.Close()

	// Remove the temp file
	err = os.Remove(tempfile.Name())
	if err != nil {
		t.Errorf("Unable to remove temp file: %v", err)
	}
}
