// The ZIP package is a plugin between the generic archive package and the standard ZIP package of the GO library.
package zip

import (
	my_archive "GoExercices/Chapter-10/Exercice-2/archive"
	"archive/zip"
	"bytes"
	"io"
	"io/fs"
	"os"
)

// ZipItem represents an item (file, folder...) of the ZIP archive.
type ZipItem struct {
	// Embedded ZIP File object of the standard GO library (used to access file data)
	file *zip.File
	// Reader to access the file data
	fileReader fs.File
	// Archive containg the File object
	zipReader *ZipReader
}

// ZipReader is a wrapper of the ZIP Reader object of the standard GO library
// It represents an ZIP archive
type ZipReader struct {
	// Embedded ZIP archive reader of the standard GO library (used to parse the archive)
	reader *zip.Reader
	// Next item (file, directory...) in the archive (used during the enumeration)
	nextItem int
	// Current ZIP item
	currentItem *ZipItem
}

// ZipArchiver is used to implement the Archiver interface required to register this package
type ZipArchiver struct{}

// init is called at package initialization
// It registers the ZIP archiver.
func init() {
	my_archive.RegisterFormat("zip", new(ZipArchiver))
}

// IsValid determines if the file is a ZIP archive
func (*ZipArchiver) IsValid(in *os.File) (bool, error) {
	// Generic ZIP pattern of a non-empty ZIP archive
	pattern := []byte{0x50, 0x4b, 0x3, 0x4}
	// Specific pattern of an empty ZIP archive
	patternEmpty := []byte{0x50, 0x4b, 0x5, 0x6}

	// Get file info
	fi, err := in.Stat()
	if err != nil {
		return false, err
	}

	// Check for minimal size (the pattern size)
	if fi.Size() < 4 {
		return false, nil
	}

	// Read the first four bytes (size of the pattern) of the file
	b := make([]byte, 4)
	l, err := in.ReadAt(b, 0)
	if err != nil {
		return false, err
	}
	if l != len(b) {
		return false, err
	}

	// Check if it is a ZIP pattern
	return bytes.Equal(b, pattern) || bytes.Equal(b, patternEmpty), nil
}

// NewReader creates an reader instance for a ZIP archive
func (*ZipArchiver) NewReader(in *os.File) (my_archive.ArchiveReader, error) {
	// Get file size (required to create the ZIP reader)
	fi, err := in.Stat()
	if err != nil {
		return nil, err
	}

	// Create the reader
	var reader ZipReader
	reader.reader, err = zip.NewReader(in, fi.Size())
	if err != nil {
		return nil, err
	}
	return &reader, nil
}

// Next returns the next file in the ZIP archive
func (zr *ZipReader) Next() (my_archive.ArchiveItem, error) {
	// Close the current file (if any)
	if zr.currentItem != nil {
		if zr.currentItem.fileReader != nil {
			zr.currentItem.fileReader.Close()
		}
		zr.currentItem = nil
	}

	// Check for EOF
	if zr.nextItem >= len(zr.reader.File) {
		return nil, io.EOF
	}

	// Create the current item
	// The file must not be opened here because this is not available for directories, sockets...
	var item ZipItem
	item.file = zr.reader.File[zr.nextItem]
	item.zipReader = zr

	// Store the current item
	zr.currentItem = &item

	// The position is updated to the next item in the archive
	zr.nextItem++

	return &item, nil
}

// GetName returns the name of the ZIP item
func (zi *ZipItem) GetName() string {
	return zi.file.Name
}

// GetSize returns the size of the ZIP item
func (zi *ZipItem) GetSize() int64 {
	return int64(zi.file.FileHeader.UncompressedSize64)
}

// GetInfo returns information about the ZIP item
func (zi *ZipItem) GetInfo() fs.FileInfo {
	return zi.file.FileHeader.FileInfo()
}

// Read gets data from the ZIP item
// This function can be called multiple times to read all data
func (zi *ZipItem) Read(buf []byte) (int, error) {
	if zi.fileReader == nil {
		var err error
		zi.fileReader, err = zi.zipReader.reader.Open(zi.file.Name)
		if err != nil {
			return 0, err
		}
	}
	return zi.fileReader.Read(buf)
}
