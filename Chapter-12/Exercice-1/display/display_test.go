package display

import (
	"testing"
)

func TestStructs(t *testing.T) {

	type Identity struct {
		FirstName string
		LastName  string
	}

	salaries := map[Identity]int{
		{"John", "Doe"}:      1000,
		{"Mary", "Smith"}:    2000,
		{"Helen", "Douglas"}: 3000,
	}

	Display("salaries", salaries)
}

func TestArrays(t *testing.T) {

	salaries := map[[2]string]int{
		{"John", "Doe"}:      1000,
		{"Mary", "Smith"}:    2000,
		{"Helen", "Douglas"}: 3000,
	}

	Display("salaries", salaries)
}
