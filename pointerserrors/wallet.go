package pointerserrors

import (
	"errors"
	"fmt"
)

type Bitcoin int // Go lets you create new types from existing ones
// 'int' will be called 'Core type' or 'Underlying type' for Bitcoin

// advantage of creting new types is that we can create methods on them
//Useful for adding more domain specific meaning to values + let you implement interfaces

// fmt package have Stringer interface with a String() method
// type Stringer interface {
// 	String() string
// }
// It will be used in fmt fucntion for formatting
// kinda like toString() in Java

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin // lowercase symbol will be pricate outside the package it's defined in
}

// "a pointer to a wallet"
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount // when you call a function or a method the arguments are copied. w without a pointer is a copy
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds // errors - Go standard package
	}
	w.balance -= amount
	return nil // return "no error"
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance // there is an impilicit dereference. You can explicitly write (*w).balance
}

// By convention you should keep your method receiver types the same for consistency (either copy or pointer)

// The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make 
// a copy of something, in which case you need to pass a reference. 
// Examples include referencing very large data structures or things where only one instance is necessary 
// (like database connection pools).

// Even tho pointers are can be confusing, you can't get segfault or memory leak in go (easily)
