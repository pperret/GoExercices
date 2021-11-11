// calc is a web calculator
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"GoExercices/Chapter-7/Exercice-16/eval"
)

// formData is the data of the calculator form
type formData struct {
	Expression string
	Variables  string
	Result     string
	Error      string
}

// tmplCalculator is the HTML template to display the calculator form
const tmplCalculator = `
<html>
	<body>
		<form action="/calc" method="POST">
			<table>
				<tr>
					<td>Expression</td>
					<td><input type="text" value="{{.Expression}}" name="expr" id="expr"/></td>
				</tr>
				<tr>
					<td>Variables</td>
					<td><input type="text" value="{{.Variables}}" name="vars" id="vars"/></td>
				</tr>
				{{if .Result}}
				<tr>
					<td>Result</td>
					<td>{{.Result}}</td>
				</tr>
				{{end}}
				{{if .Error}}
				<tr>
					<td>Error</td>
					<td>{{.Error}}</td>
				</tr>
				{{end}}
			</table>
			<input type="submit" value="Calc">
		</form>
	</body>
</html>
`

// formCalculator is the compiled version of the calculator template
var formCalculator = template.Must(template.New("calculator").Parse(tmplCalculator))

// displayError displays an error in the calculator form
func displayError(w http.ResponseWriter, expression, variables, errorFormat string, errorArgs ...interface{}) {
	data := formData{Expression: expression, Variables: variables, Error: fmt.Sprintf(errorFormat, errorArgs...)}
	if err := formCalculator.Execute(w, data); err != nil {
		log.Fatalf("Unable to display the calculator form: err=%v", err)
	}
}

// home displays the empty form of the calculator
func home(w http.ResponseWriter, r *http.Request) {
	data := formData{}
	if err := formCalculator.Execute(w, data); err != nil {
		log.Fatalf("Unable to display the calculator form: err=%v", err)
	}
}

// calc compute the result
func calc(w http.ResponseWriter, r *http.Request) {
	// Parse the form
	if err := r.ParseForm(); err != nil {
		displayError(w, "", "", "unable to parse form: %v", err)
		return
	}

	// Get the form data
	expression := r.Form.Get("expr")
	variables := r.Form.Get("vars")

	// Check that expression is not empty (variables may be empty)
	if expression == "" {
		displayError(w, expression, variables, "empty expression")
		return
	}

	// Parse the expression
	expr, err := eval.Parse(expression)
	if err != nil {
		displayError(w, expression, variables, "invalid expression: %v", err)
		return
	}

	// Check the expression and get the list of required variables
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		displayError(w, expression, variables, "invalid expression: %v", err)
		return
	}

	// Parse the variables
	env := eval.Env{}
	vl := strings.FieldsFunc(variables, func(r rune) bool { return r == ',' || r == ';' })
	for _, v := range vl {
		v = strings.TrimSpace(v)
		v2 := strings.Split(v, "=")
		if len(v2) != 2 {
			displayError(w, expression, variables, "invalid variable %s", v)
			return
		}
		name := strings.TrimSpace(v2[0])
		value := strings.TrimSpace(v2[1])
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			displayError(w, expression, variables, "invalid value for %s", name)
			return
		}
		env[eval.Var(name)] = f
	}

	// Check that required variables are available
	for v := range vars {
		if _, ok := env[v]; !ok {
			displayError(w, expression, variables, "%s is not set", v)
			return
		}
	}

	// Display the result
	data := formData{Expression: expression, Variables: variables, Result: fmt.Sprintf("%g", expr.Eval(env))}
	if err = formCalculator.Execute(w, data); err != nil {
		log.Fatalf("Unable to display the items list: err=%v", err)
	}
}

// main is the entry point of the program
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
