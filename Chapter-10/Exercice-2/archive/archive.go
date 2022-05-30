package archive

import (
	"fmt"
	"io/fs"
	"os"
)

// ArchiveItem is the interface to implement by an archive item
type ArchiveItem interface {
	// GetName returns the name of the archive item
	GetName() string
	// GetSize returns the data size of the archive item
	GetSize() int64
	// GetInfo returns information about the archive item
	GetInfo() fs.FileInfo
	// Read gets data from the archive item
	Read([]byte) (int, error)
}

// ArchiveReader is the interface to implement by an archive reader
type ArchiveReader interface {
	// Next returns the next item in the archive
	Next() (ArchiveItem, error)
}

// Archiver is the interface to implement by each archiver
type Archiver interface {
	// IsValid determines if the archiver is able to manage the archive
	IsValid(*os.File) (bool, error)
	// NewReader creates an reader instance of the archiver
	NewReader(*os.File) (ArchiveReader, error)
}

// ArchiverEntry is registered archiver (in the list)
type ArchiverEntry struct {
	// name is the name of the archiver
	name string
	// archiver is the interface to access the archiver
	archiver Archiver
}

// archiverList is the list of registered archivers
var archiverList []ArchiverEntry

// RegisterFormat is used to register an archiver
// name is the archiver name
func RegisterFormat(name string, archiver Archiver) {
	archiverList = append(archiverList, ArchiverEntry{name, archiver})
}

// NewReader is intended to create a reader instance according to the archive type
// The reader (corresponding to the archive format) is automatically determined
// The function returns the created reader, the name of the archiver and an error (or nil)
func NewReader(in *os.File) (ArchiveReader, string, error) {
	// Loop on the registered archivers
	for _, entry := range archiverList {
		// Check if the archiver is able to manage the archive
		valid, err := entry.archiver.IsValid(in)
		if err != nil {
			return nil, "", err
		}
		// If the current archiver is able to manager the archive, a reader instance is created
		if valid {
			reader, err := entry.archiver.NewReader(in)
			if err != nil {
				return nil, entry.name, err
			}
			return reader, entry.name, nil
		}
	}
	// There is no archiver to take into account the archive
	return nil, "", fmt.Errorf("no corresponding archive reader")
}
