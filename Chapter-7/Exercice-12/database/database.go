// database is a sample HTTP server managing a memory database
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// main is the entry point of the program
func main() {
	var db database
	db.items = make(map[string]dollars)
	db.items["shoes"] = 50
	db.items["socks"] = 5
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// dollars is the price of a database item
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// database is a memory database based on map
type database struct {
	items map[string]dollars
	mutex sync.Mutex
}

// tmplItems is the HTML template to display the items list
const tmplItems = `
<html>
	<body>
		<table>
			<tr style='textalign:left'>
				<th>Item</th>
				<th>Price</th>
			</tr>
			{{range $key, $value := .}}
			<tr>
				<td>{{$key}}</td>
				<td>{{$value}}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>
`

// reportItemList is the compiled version of the template to display the items list
var reportItemList = template.Must(template.New("itemList").Parse(tmplItems))

// list lists the database content
// URL: /list
func (db *database) list(w http.ResponseWriter, req *http.Request) {
	// Lock the items list
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Print the items
	err := reportItemList.Execute(w, db.items)
	if err != nil {
		log.Fatalf("Unable to display the items list: err=%v", err)
	}
}

// price returns the price of a database item
// URL: /price?item={name}
func (db *database) price(w http.ResponseWriter, req *http.Request) {
	// Get the "item" parameter
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item is not specified: %q\n", item)
		return
	}

	// Lock the items list
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Look for the item price
	price, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	// Display the price
	fmt.Fprintf(w, "%s\n", price)
}

// create appends an entry into the database
// URL: /create?item={name}&price={price}
func (db *database) create(w http.ResponseWriter, req *http.Request) {
	// Get URL parameters
	values := req.URL.Query()

	// Get the "item" parameter
	item := values.Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item is not specified\n")
		return
	}

	// Get the "price" parameter
	priceString := values.Get("price")
	if priceString == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price is not specified\n")
		return
	}

	// Convert price in float
	price, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price is invalid: %q\n", priceString)
		return
	}

	// Lock the items list
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Check the item not already exists
	_, ok := db.items[item]
	if ok {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	// Append the item to the database
	db.items[item] = dollars(price)

	fmt.Fprintf(w, "Item created\n")
}

// read gets a database entry
// URL: /read?item={name}
func (db *database) read(w http.ResponseWriter, req *http.Request) {
	// Get the "item" parameter
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item is not specified\n")
		return
	}

	// Lock the items list
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Look for the item price
	price, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	// Display the item
	fmt.Fprintf(w, "%s -> %s\n", item, price)
}

// update updates the price of a database item
// URL: /update?item={name}&price={price}
func (db *database) update(w http.ResponseWriter, req *http.Request) {
	// Get URL parameters
	values := req.URL.Query()

	// Get the "item" parameter
	item := values.Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item is not specified\n")
		return
	}

	// Get the "price" parameter
	priceString := values.Get("price")
	if priceString == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price is not specified\n")
		return
	}

	// Convert price in float
	price, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price is invalid: %q\n", priceString)
		return
	}

	// Lock the items list
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Check if the item exists
	_, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item does not exist: %q\n", item)
		return
	}

	// Update the item to the database
	db.items[item] = dollars(price)

	fmt.Fprintf(w, "Item updated\n")
}

// delete deletes a database item
// URL: /delete?item={name}
func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	// Get the "item" parameter
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item is not specified\n")
		return
	}

	// Lock the items list
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Check if the item exists
	_, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	// Delete the item
	delete(db.items, item)

	fmt.Fprintf(w, "Item deleted\n")
}
