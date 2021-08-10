package main

import (
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

// URL is the HTML page to test
const URL = "http://golang.org"

// TestPrettyPrint1 tests the prettyPrint function using a pipe for stdout
func TestPrettyPrint1(t *testing.T) {

	// Redirect Stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute the test
	go func() {
		if err := prettyPrint(URL); err != nil {
			t.Fatalf("Error while formatting %v", err)
		}
		w.Close()
	}()

	// Parse the result
	_, err := html.Parse(r)
	if err != nil {
		t.Fatalf("Error while parsing result: %v", err)
	}

	r.Close()
	os.Stdout = old
}

// TestPrettyPrint2 tests the prettyPrint function using a temporay file for stdout
func TestPrettyPrint2(t *testing.T) {

	// Redirect Stdout to a temp file
	old := os.Stdout
	tempfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("Unable to open temp file: %v", err)
	}
	os.Stdout = tempfile

	// Execute the test
	if err := prettyPrint(URL); err != nil {
		t.Fatalf("Error while formatting %v", err)
	}

	// Parse the result
	_, err = html.Parse(tempfile)
	if err != nil {
		t.Fatalf("Error while parsing result: %v", err)
	}

	// Restore stdout
	os.Stdout = old
	tempfile.Close()

	// Remove the temp file
	err = os.Remove(tempfile.Name())
	if err != nil {
		t.Errorf("Unable to remove temp file: %v", err)
	}
}
