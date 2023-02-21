package bzip_test

import (
	"bytes"
	"compress/bzip2" // reader
	"io"
	"sync"
	"testing"

	"GoExercices/Chapter-13/Exercice-3/bzip" // writer
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}

func TestBzip2Sync(t *testing.T) {
	const nbGoRoutines = 10
	var compressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// Create several goruntines in order to compress data concurently
	var wg sync.WaitGroup
	for n := 0; n < nbGoRoutines; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Write a repetitive message in a million pieces
			for i := 0; i < 1000000; i++ {
				io.WriteString(w, "hello")
			}
		}()
	}
	wg.Wait()

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 2296; got != want {
		t.Errorf("hellos compressed to %d bytes, want %d", got, want)
	}
}
