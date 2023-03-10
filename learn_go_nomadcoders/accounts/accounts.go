package accounts

import "errors"

// Account struct
type Account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("can't withdraw")
// NewAccount creates a new account
func NewAccount(owner string) *Account {
	account := Account{owner : owner, balance : 0}
	return &account 
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int){
	a.balance += amount
}

//Balance of your account
func (a Account) Balance() int{
	return a.balance
}

//Withdraw x amount on your account
func (a *Account) WithBalance(amount int) error {

	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

//Change Owner of the account
func (a *Account) ChangeOwner(newOwner string){
	a.owner = newOwner
}

//Owner of the account
func (a Account) Owner() string{
	return a.owner
}