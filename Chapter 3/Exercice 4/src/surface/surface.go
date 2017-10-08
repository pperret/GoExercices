// Server2 is a minimal "echo" and counter server.
package main
import (
	"strconv"
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	cells = 100 // number of grid cells
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler generating the graphic
func handler(w http.ResponseWriter, r *http.Request) {

	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}

	// Get height
	height := 320 // Default value
	ht := r.Form["height"]
	if ht != nil {
		h, err := strconv.ParseUint(ht[0], 10, 64)
		if err == nil {
			height = int(h)
		}
	}

	// Get width
	width := 600 // Default value
	wt := r.Form["width"]
	if wt != nil {
		w, err := strconv.ParseUint(wt[0], 10, 64)
		if err == nil {
			width = int(w)
		}
	}
	
	// Get color
	color := "grey" // Default value
	ct := r.Form["color"]
	if ct != nil {
		color = ct[0]
	}
	
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' " +
	"style='stroke: %s; fill: white; strokewidth:0.7' " +
	"width='%d' height='%d'>\n", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok_a := corner(width, height, i + 1, j)
			bx, by, ok_b := corner(width, height, i, j)
			cx, cy, ok_c := corner(width, height, i, j + 1)
			dx, dy, ok_d := corner(width, height, i + 1, j + 1)
			if ok_a && ok_b && ok_c && ok_d {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(width, height, i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := float64(height) * 0.4 // pixels per z unit	

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width) / 2 + (x - y) * cos30 * xyscale
	sy := float64(height) / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}