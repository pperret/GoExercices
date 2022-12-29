// Web server delivering Lissajous figures
package main

import (
	"GoExercices/Chapter-12/Exercice-12/params"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
)

// main is the entry point of the program
func main() {
	params.CheckMap["email"] = checkEMail
	params.CheckMap["credit_card"] = checkCreditCard
	params.CheckMap["zip"] = checkZIPCode

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler for the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email      string `http:"email" check:"email"`
		CreditCard string `http:"cc" check:"credit_card"`
		ZIPCode    int    `http:"zip" check:"zip"`
	}

	// Parse request parameters
	if err := params.Unpack(r, &data); err != nil {
		fmt.Fprintf(w, "Bad parameters: %v", err)
		return
	}

	// Display resulting values
	if data.Email != "" {
		fmt.Fprintf(w, "<p>Email address: %s</p>", data.Email)
	}
	if data.CreditCard != "" {
		fmt.Fprintf(w, "<p>Credit card number: %s</p>", data.CreditCard)
	}
	if data.ZIPCode != 0 {
		fmt.Fprintf(w, "<p>ZIP code: %d</p>", data.ZIPCode)
	}
}

// checkEmail validates an email address
func checkEMail(value string) error {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

// checkLuhn validate the Luhn key
func checkLuhn(value string) error {
	sum := 0
	for i, c := range value {
		idx := strings.IndexRune("0123456789", c)
		if idx == -1 {
			return fmt.Errorf("invalid character in credit card number")
		}
		if i%2 == (len(value)+1)%2 {
			sum = sum + idx
		} else {
			if 2*idx > 10 {
				sum = sum + 2*idx - 9
			} else {
				sum = sum + 2*idx
			}
		}
	}
	if sum%10 != 0 {
		return fmt.Errorf("invalid credit card number (bad Luhn key)")
	}
	return nil
}

// checkCreditCard validates a credit card number
func checkCreditCard(value string) error {
	// Check number length
	if len(value) < 2 {
		return fmt.Errorf("credit card number is too short")
	}
	switch value[0] {
	case '3':
		switch value[1] {
		case '4', '7': // AMEX
			if len(value) != 15 {
				return fmt.Errorf("invalid length of credit card number")
			}
		case '0', '6', '8': // Diners
			if len(value) != 14 {
				return fmt.Errorf("invalid length of credit card number")
			}
		}
	case '4': // Visa
		if len(value) != 13 && len(value) != 16 {
			return fmt.Errorf("invalid length of credit card number")
		}
	case '5': // MasterCard
		if len(value) != 16 {
			return fmt.Errorf("invalid length of credit card number")
		}
	case '6': // Discover
		if len(value) != 16 {
			return fmt.Errorf("invalid length of credit card number")
		}
	}

	// Check the Luhn key
	return checkLuhn(value)
}

// checkZIPCode validates a US ZIP code
func checkZIPCode(value string) error {
	zip, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid US ZIP code (not integer)")
	}
	if zip == 0 || zip > 99999 {
		return fmt.Errorf("invalid US ZIP code")
	}
	return nil
}
