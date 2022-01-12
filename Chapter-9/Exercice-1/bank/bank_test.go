package main

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		ok := Withdraw(1000)
		if ok {
			t.Errorf("Alice: balance allows a withdrawal of 1000")
		}
		Deposit(200)
		ok = Withdraw(200)
		if !ok {
			t.Errorf("Alice: balance reject a withdrawal of 200")
		}
		fmt.Println("Alice=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		ok := Withdraw(200)
		if ok {
			t.Errorf("Charles: balance allows a withdrawal of 200")
		}
		Deposit(100)
		ok = Withdraw(100)
		if !ok {
			t.Errorf("Charles: balance rejects a withdrawal of 100")
		}
		fmt.Println("Bob=", Balance())
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
