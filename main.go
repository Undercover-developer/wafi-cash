package main

/*
	All the code was put in the same file for simplicity and easy inspection.
*/
import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type User struct {
	Id   int
	Name string
}
type Account struct {
	UserId  int
	Balance float64
}

type Transaction struct {
	FromUserId int
	ToUserId   int
	Amount     float64
}

func AddUser(name string) User {
	id := rand.Intn(100)
	newUser := User{
		Id:   id,
		Name: name,
	}
	AddAccount(newUser)
	users = append(users, newUser)
	return newUser
}

func AddAccount(user User) {
	accounts = append(accounts, Account{
		UserId:  user.Id,
		Balance: 0,
	})
}

func Deposit(user User, amount float64) {
	for i, account := range accounts {
		if account.UserId == user.Id {
			accounts[i].Balance += amount
		}
	}
}

func GetBalance(user User) float64 {
	for _, account := range accounts {
		if account.UserId == user.Id {
			return account.Balance
		}
	}
	return 0
}

func Transfer(fromUserId int, toUserId int, amount float64) {
	for i, account := range accounts {
		if account.UserId == fromUserId {
			if account.Balance >= amount {
				accounts[i].Balance -= amount
			}
			return
		}
		if account.UserId == toUserId {
			accounts[i].Balance += amount
		}
	}
}

// unit test for app functions
func TestAddUser(t *testing.T) {
	user := AddUser("User A")
	if user.Name != "User A" {
		t.Errorf("User name is not correct")
	}
}

func TestDeposit(t *testing.T) {
	user := AddUser("User A")
	Deposit(user, 10)
	if accounts[0].Balance != 10 {
		t.Errorf("Deposit is not correct")
	}
}

func TestAddAccount(t *testing.T) {
	user := AddUser("User A")
	AddAccount(user)
	if len(accounts) != 1 {
		t.Errorf("Account is not added")
	}
}

func TestGetBalance(t *testing.T) {
	user := AddUser("User A")
	Deposit(user, 10)
	if GetBalance(user) != 10 {
		t.Errorf("GetBalance is not correct")
	}
}

func TestTransfer(t *testing.T) {
	userA := AddUser("User A")
	Deposit(userA, 10)
	userB := AddUser("User B")
	Deposit(userB, 20)
	Transfer(userB.Id, userA.Id, 15)
	if GetBalance(userA) != 25 {
		t.Errorf("Transfer is not correct")
	}
}

func runTests() {
	t := new(testing.T)
	TestAddUser(t)
	TestAddAccount(t)
	TestDeposit(t)
	TestGetBalance(t)
	TestTransfer(t)
}

var users []User
var accounts []Account

func main() {
	trxCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Add usersA
	userA := AddUser("User A")
	//deposit 10 dollars
	Deposit(userA, 10)
	// Add usersB
	userB := AddUser("User B")
	//deposit 20 dollars
	Deposit(userB, 20)
	// User B sends 15 dollars to User A
	trxChannel := make(chan Transaction) // to simulate transaction queue
	trx := Transaction{
		FromUserId: userB.Id,
		ToUserId:   userA.Id,
		Amount:     15,
	}
	//schedule transaction queue handler
	go func() {
		for {
			select {
			case <-trxCtx.Done():
				return
			case trx := <-trxChannel:
				Transfer(trx.FromUserId, trx.ToUserId, trx.Amount)
			}
		}
	}()
	trxChannel <- trx

	// User A checks their balance and has 25 dollars
	fmt.Println("User A balance is: ", GetBalance(userA))
	// User B checks their balance and has 5 dollars
	fmt.Println("User B balance is: ", GetBalance(userB))
	// User A transfers 25 dollars from their account
	trx = Transaction{
		FromUserId: userA.Id,
		ToUserId:   0,
		Amount:     25,
	}
	trxChannel <- trx
	time.Sleep(time.Second)
	// User A checks their balance and has 0 dollars
	fmt.Println("User A balance is: ", GetBalance(userA))
}
