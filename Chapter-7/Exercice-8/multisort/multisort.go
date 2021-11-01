// stablesort is an implementation of a multi-columns sort
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

// length converts a string as a time.Duration
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// Track is the data of a music track
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// tracks is the initial content of the track list
var tracks = []*Track{
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Go", "Delilah2", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Delilah2", "From the Roots Up", 2013, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
}

// printTracks displays the content of the track list
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

// lessFunction is the prototype of Less functions used to sort column elements
type lessFunction func(x, y *Track) bool

// lessXxxxx functions are used to compare elements of track columns
func lessTitle(x, y *Track) bool  { return strings.Compare(x.Title, y.Title) < 0 }
func lessArtist(x, y *Track) bool { return strings.Compare(x.Artist, y.Artist) < 0 }
func lessAlbum(x, y *Track) bool  { return strings.Compare(x.Album, y.Album) < 0 }
func lessYear(x, y *Track) bool   { return x.Year < y.Year }
func lessLength(x, y *Track) bool { return x.Length < y.Length }

// multiSort is an object used for columns sort
type multiSort struct {
	tracks []*Track     // Track list
	less   lessFunction // Compare function to be used
}

// Implementation of the sort.Interface functions
func (ms multiSort) Len() int           { return len(ms.tracks) }
func (ms multiSort) Less(i, j int) bool { return ms.less(ms.tracks[i], ms.tracks[j]) }
func (ms multiSort) Swap(i, j int)      { ms.tracks[i], ms.tracks[j] = ms.tracks[j], ms.tracks[i] }

// clickByXxxxx functions perform a sort as if the user had clicked on a column header
func clickByTitle(tracks []*Track)  { sort.Sort(multiSort{tracks, lessTitle}) }
func clickByArtist(tracks []*Track) { sort.Sort(multiSort{tracks, lessArtist}) }
func clickByAlbum(tracks []*Track)  { sort.Sort(multiSort{tracks, lessAlbum}) }
func clickByYear(tracks []*Track)   { sort.Sort(multiSort{tracks, lessYear}) }
func clickByLength(tracks []*Track) { sort.Sort(multiSort{tracks, lessLength}) }

// main is the entry point of the program
func main() {

	// Print the initial content of the track list
	printTracks(tracks)

	// Perform sorts as if column headers have been clicked
	clickByLength(tracks)
	clickByYear(tracks)
	clickByArtist(tracks)
	clickByAlbum(tracks)
	clickByTitle(tracks)

	// Print the final content of the track list
	printTracks(tracks)
}
