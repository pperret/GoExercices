// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strings"
)

// Check populates the variable list with the current variable
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

// Check does not perform any operation for literals
func (literal) Check(vars map[Var]bool) error {
	return nil
}

// Check verifies the unitary operator then checks the operand recursively
func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

// Check verifies the binary operator then checks the operands recursively
func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

// Check verifies the function name and the arguments count then checks the arguments recursively
func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

// numParams is the number of arguments for each supported function
var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

// Check verifies the operation name and the arguments count then checks the arguments recursively
func (s set) Check(vars map[Var]bool) error {
	_, ok := setOperations[s.op]
	if !ok {
		return fmt.Errorf("unknown operation %q", s.op)
	}
	for _, arg := range s.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var setOperations = map[string]bool{"min": true, "max": true, "count": true, "sum": true, "avg": true}
