// The TAR package is a plugin between the generic archive package and the standard TAR package of the GO library.
package tar

import (
	my_archive "GoExercices/Chapter-10/Exercice-2/archive"
	"archive/tar"
	"bytes"
	"io/fs"
	"os"
)

// TarItem represents a item (file, folder...) of the TAR archive.
type TarItem struct {
	// Embedded TAR header object of the standard GO library
	header *tar.Header
	// The TAR reader is used to access file data
	reader *TarReader
}

// TarReader is a wrapper of the TAR Reader object of the standard GO library
// It represents a TAR archive
type TarReader struct {
	reader *tar.Reader
}

// TarArchiver is used to implement the Archiver interface required to register this package
type TarArchiver struct{}

// init is called at package initialization
// It registers the TAR archiver.
func init() {
	my_archive.RegisterFormat("tar", new(TarArchiver))
}

// IsValid determines if the file is a TAR archive
func (*TarArchiver) IsValid(in *os.File) (bool, error) {
	// Pattern of a TAR archive
	pattern := []byte{'u', 's', 't', 'a', 'r'}

	// Get file info
	fi, err := in.Stat()
	if err != nil {
		return false, err
	}

	// Check for minimal size (tar header size)
	if fi.Size() < 512 {
		return false, nil
	}

	// Read 5 bytes (size of the pattern) located at offset 257
	b := make([]byte, 5)
	l, err := in.ReadAt(b, 257)
	if err != nil {
		return false, err
	}
	if l != len(b) {
		return false, err
	}

	// Check if it is the TAR pattern
	return bytes.Equal(b, pattern), nil
}

// NewReader creates an reader instance for a TAR archive
func (*TarArchiver) NewReader(in *os.File) (my_archive.ArchiveReader, error) {
	var reader TarReader
	reader.reader = tar.NewReader(in)
	return &reader, nil
}

// Next returns the next file in the TAR archive
func (tr *TarReader) Next() (my_archive.ArchiveItem, error) {
	// Get the header of the next item in the archive
	header, err := tr.reader.Next()
	if err != nil {
		return nil, err
	}

	// Create the instance of the TarItem
	var item TarItem
	item.header = header
	item.reader = tr
	return &item, nil
}

// GetName returns the name of the TAR item
func (ti *TarItem) GetName() string {
	return ti.header.Name
}

// GetSize returns the size of the TAR item
func (ti *TarItem) GetSize() int64 {
	return ti.header.Size
}

// GetInfo returns information about the TAR itel
func (ti *TarItem) GetInfo() fs.FileInfo {
	return ti.header.FileInfo()
}

// Read gets data from the TAR item
// This function can be called multiple times to read all data
func (ti *TarItem) Read(buf []byte) (int, error) {
	return ti.reader.reader.Read(buf)
}
