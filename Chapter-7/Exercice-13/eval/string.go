package eval

import (
	"bytes"
	"fmt"
)

// String generates the string format of the variable
func (v Var) String() string {
	return string(v)
}

// String generates the string format of the literal
func (l literal) String() string {
	return fmt.Sprintf("%f", float64(l))
}

// String generates the string format of the unitary expression
func (u unary) String() string {
	return string(u.op) + u.x.String()
}

// String generates the string format of the binary expression
func (b binary) String() string {
	return "(" + b.x.String() + " " + string(b.op) + " " + b.y.String() + ")"
}

// String generates the string format of the function call
func (c call) String() string {
	var sb bytes.Buffer
	sb.WriteString(c.fn + "(")
	for i, arg := range c.args {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(arg.String())
	}
	sb.WriteString(")")
	return sb.String()
}
