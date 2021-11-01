// stablesort is an implementation of a multi-columns sort
package main

import (
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"
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

// tmplTracks is the HTML template to display the track list
const tmplTracks = `
<html>
	<head>
		<meta charset="utf-8"/>
		<style media="screen" type="text/css">
			table {
				border-collapse: collapse;
				border-spacing: 0px;
			}
			table, th, td {
				padding: 5px;
				border: 1px solid black;
			}
		</style>
	</head>
	<body>
		<table>
			<thead>
				<tr style='textalign:left'>
					<th><a href='/?sortBy=title'>Title</a></th>
					<th><a href='/?sortBy=artist'>Artist</a></th>
					<th><a href='/?sortBy=album'>Album</a></th>
					<th><a href='/?sortBy=year'>Year</a></th>
					<th><a href='/?sortBy=length'>Length</a></th>
				</tr>
			</thead>
			<tbody>
				{{range .}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Artist}}</td>
					<td>{{.Album}}</td>
					<td>{{.Year}}</td>
					<td>{{.Length}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`

// reportTrackList is the compiled version of the template to display the track list
var reportTrackList = template.Must(template.New("trackList").Parse(tmplTracks))

// displayTracks displays the content of the track list as HTML
func displayTracks(w http.ResponseWriter, tracks []*Track) {

	err := reportTrackList.Execute(w, tracks)
	if err != nil {
		log.Fatalf("Unable to display track list: err=%v", err)
	}
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

// handler sorts the track list then displays the updated content
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("sortBy") {
	case "title":
		clickByTitle(tracks)
	case "artist":
		clickByArtist(tracks)
	case "album":
		clickByAlbum(tracks)
	case "year":
		clickByYear(tracks)
	case "length":
		clickByLength(tracks)
	}
	displayTracks(w, tracks)
}

// main is the entry point of the program
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
