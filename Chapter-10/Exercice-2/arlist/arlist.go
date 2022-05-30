// The arlist command reads an archive (in tar, zip... format)
// and lists its content
package main

import (
	my_archive "GoExercices/Chapter-10/Exercice-2/archive"
	_ "GoExercices/Chapter-10/Exercice-2/archive/tar"
	_ "GoExercices/Chapter-10/Exercice-2/archive/zip"
	"fmt"
	"io"
	"os"
)

// main is the entry point of the program
func main() {

	// Check command arguments
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(2)
	}

	// Open the input archive file
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open input file: %v", err)
		os.Exit(3)
	}
	defer f.Close()

	// List the archive items
	if err := list(f); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

// list lists the archive (tar or zip) items
func list(in *os.File) error {

	// Create the reader corresponding to the archive type
	reader, kind, err := my_archive.NewReader(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	// List the archive items
	for {
		file, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to get next archive's file: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "%v %v %12d %s\n", file.GetInfo().Mode(), file.GetInfo().ModTime(), file.GetSize(), file.GetName())
	}
	return nil
}
