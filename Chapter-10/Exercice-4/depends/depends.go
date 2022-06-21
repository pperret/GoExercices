// depends is a tool identifying workspace packages depending transitively of packages specified as parameters
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// usage displays the program usage
func usage(arg0 string) {
	fmt.Fprintf(os.Stderr, "Usage: %s <package> [packages...]\n", arg0)
}

// Pack contains data (only those needed by the program) about a package.
type Pack struct {
	Name       string   // Package name
	ImportPath string   // Package import path
	Deps       []string // Package dependencies
}

// list gets information about a list of packages whose names are provided as parameters
func list(names ...string) ([]Pack, error) {
	// Use "go list" command to get dependencies of packages in JSON format
	var args []string
	args = append(args, "list")
	args = append(args, "-json")
	args = append(args, names...)
	cmd := exec.Command("go", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Analyze dependencies using a JSON decoder
	reader := bytes.NewReader(out)
	dec := json.NewDecoder(reader)
	var packs []Pack
	for {
		var pack Pack
		err := dec.Decode(&pack)
		if err == io.EOF {
			return packs, nil

		} else if err != nil {
			return nil, err
		}
		packs = append(packs, pack)
	}
}

// getImportPaths provides import path of a packages list
// It returns a map indexed by the import paths
func getImportPaths(names ...string) (map[string]string, error) {

	// Get packges information
	packs, err := list(names...)
	if err != nil {
		return nil, err
	}

	// Extract import paths of packages
	importPaths := make(map[string]string)
	for _, pack := range packs {
		importPaths[pack.ImportPath] = pack.Name
	}
	return importPaths, nil

}

// getDependencies provides dependencies of a packages list
// It returns a map indexed by the package names
func getDependencies(names ...string) (map[string][]string, error) {

	// Get packages information
	packs, err := list(names...)
	if err != nil {
		return nil, err
	}

	// Extract packages dependenciess
	deps := make(map[string][]string)
	for _, pack := range packs {
		deps[pack.Name] = append(deps[pack.Name], pack.Deps...)
	}
	return deps, nil
}

// main is the entry point of the program
func main() {
	if len(os.Args) < 2 {
		usage(os.Args[0])
		os.Exit(1)
	}

	// Get import paths of packages
	packagesImportPaths, err := getImportPaths(os.Args[1:]...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get packages import paths: %v\n", err)
		os.Exit(2)
	}

	// Get dependencies of all the packages in current workspace
	workspaceDependencies, err := getDependencies("...")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get workspace dependencies: %v\n", err)
		os.Exit(2)
	}

	// Select packages in the current workspace depending of those provided as parameters
	depending := make(map[string]int)
	for name, deps := range workspaceDependencies {
		for _, dep := range deps {
			_, ok := packagesImportPaths[dep]
			if ok {
				depending[name]++
				break
			}
		}
	}

	// Display the selected packages
	for n := range depending {
		fmt.Println(n)
	}
}
