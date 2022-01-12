// bank provides a concurrency-safe bank with one account.
package main

type WithdrawMessage struct {
	amount int
	reply  chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan WithdrawMessage)

// Deposit deposites an amount in the bank account
func Deposit(amount int) { deposits <- amount }

// Balance retrieves the current balance of the bank account
func Balance() int { return <-balances }

// Withdraw withdraws an amount from the bank account
func Withdraw(amount int) bool {
	reply := make(chan bool)
	defer close(reply)
	withdraws <- WithdrawMessage{amount, reply}
	return <-reply
}

// teller is a monitor serializing accesses to balance
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdraw := <-withdraws:
			if balance >= withdraw.amount {
				balance -= withdraw.amount
				withdraw.reply <- true
			} else {
				withdraw.reply <- false
			}
		case balances <- balance:
		}
	}
}

// init starts the balance monitoring
func init() {
	go teller() // start the monitor goroutine
}
