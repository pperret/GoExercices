// The surface program plots the 3-D surface of a user-provided function.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"GoExercices/Chapter-7/Exercice-15/eval"
)

// main is the entry point of the program
func main() {

	reader := bufio.NewReader(os.Stdin)

	// Get the expression
	var formula string
	for {
		fmt.Print("Enter expression: ")
		if s, err := reader.ReadString('\n'); err == nil {
			if s = strings.TrimSpace(s); s != "" {
				formula = s
				break
			}
		}
	}

	// Parse the expression
	expr, err := eval.Parse(formula)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid expression: %v\n", err)
		os.Exit(2)
	}

	// Check the expression and get the list of variables
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Fprintf(os.Stderr, "invalid expression: %v\n", err)
		os.Exit(2)
	}

	// Get the values of variables
	env := eval.Env{}
	for v := range vars {
		for {
			fmt.Fprintf(os.Stdin, "Enter %s: ", string(v))
			if s, err := reader.ReadString('\n'); err == nil {
				if s = strings.TrimSpace(s); s != "" {
					if f, err := strconv.ParseFloat(s, 64); err == nil {
						env[v] = f
					}
					fmt.Fprintf(os.Stderr, "invalid float: %s\n", s)
				}
			}
		}
	}

	// Compute the expression using the environment variables
	res := expr.Eval(env)

	// Print the result
	fmt.Printf("%s => %g\n", formula, res)
}
