package main

import "errors"

var (
	ErrNegativeBalance = errors.New("new balance can't be negative")
)

type Account struct {
	balance float64
	owner   string
}

func (acc *Account) SetBalance(num float64) error {
	if num >= 0 {
		acc.balance = num
		return nil
	} else {
		return ErrNegativeBalance
	}
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(num float64) error {
	if num >= 0 {
		acc.balance += num
	} else {
		return ErrNegativeBalance
	}
	return nil
}

func (acc *Account) Withdraw(num float64) error {
	if acc.balance-num >= 0 && num >= 0 {
		acc.balance -= num
	} else {
		return ErrNegativeBalance
	}
	return nil
}

func NewAccount(name string) Account {
	return Account{0, name}
}
