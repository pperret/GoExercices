// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

// Env is the list of variables (name/value)
type Env map[Var]float64

// Eval returns the value of the variable
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Eval returns the value of the literal
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// Eval returns the value of the unitary expression
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// Eval returns the value of the binary expression
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

// Eval returns the value of the function call
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

// Eval returns the value of the set operation
func (s set) Eval(env Env) float64 {
	switch s.op {
	case "min":
		var min float64
		if len(s.args) > 0 {
			min = s.args[0].Eval(env)
			for _, expr := range s.args[1:] {
				v := expr.Eval(env)
				if v < min {
					min = v
				}
			}
		}
		return min
	case "max":
		var max float64
		if len(s.args) > 0 {
			max = s.args[0].Eval(env)
			for _, expr := range s.args[1:] {
				v := expr.Eval(env)
				if v > max {
					max = v
				}
			}
		}
		return max
	case "sum":
		var sum float64
		for _, expr := range s.args {
			sum += expr.Eval(env)
		}
		return sum
	case "avg":
		var sum float64
		for _, expr := range s.args {
			sum += expr.Eval(env)
		}
		return sum / float64(len(s.args))
	case "count":
		return float64(len(s.args))
	}
	panic(fmt.Sprintf("unsupported set operation: %s", s.op))
}
