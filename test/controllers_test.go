package test

import (
	"fmt"
	"testing"
	"wafi-cash/controllers"
	"wafi-cash/models"
)

// unit test for app functions
func TestAddAccount(t *testing.T) {
	user := controllers.AddUser("User B")
	t.Log(user)
	if len(controllers.GetAccounts()) != 1 {
		t.Errorf("Account is not added")
	}
}

func TestAddUser(t *testing.T) {
	user := controllers.AddUser("User A")
	if user.Name != "User A" {
		t.Errorf("User name is not correct")
	}
}

func TestDeposit(t *testing.T) {
	user := controllers.AddUser("User A")
	controllers.Deposit(user, 10, models.USD)
	if controllers.GetBalance(user, models.USD) != 10 {
		t.Errorf("Deposit is not correct")
	}
}

func TestGetBalance(t *testing.T) {
	user := controllers.AddUser("User A")
	controllers.Deposit(user, 10, models.USD)
	if controllers.GetBalance(user, models.USD) != 10 {
		t.Errorf("GetBalance is not correct")
	}
}

func TestTransfer(t *testing.T) {
	userA := controllers.AddUser("User A")
	controllers.Deposit(userA, 10, models.USD)
	userB := controllers.AddUser("User B")
	controllers.Deposit(userB, 20, models.USD)
	controllers.Transfer(userB.Id, userA.Id, 15, models.USD)
	if controllers.GetBalance(userA, models.USD) != 25 {
		t.Errorf("Transfer is not correct")
	}
}

/*
userA deposit 2USD
then 400 NGN,
then 1GBP,
then userA transfers 3USD to userB.
userA's balance after the transfer should then become 0USD, 400NGN and 0.16GBP
*/

func TestTransfer2(t *testing.T) {
	userA := controllers.AddUser("User A")
	userB := controllers.AddUser("User B")
	controllers.Deposit(userA, 2, models.USD)
	controllers.Deposit(userA, 400, models.NGN)
	controllers.Deposit(userA, 1, models.GBP)
	controllers.Transfer(userA.Id, userB.Id, 3, models.USD)
	fmt.Println(controllers.GetBalance(userA, models.USD))
	fmt.Println(controllers.GetBalance(userA, models.NGN))
	fmt.Println(controllers.GetBalance(userA, models.GBP))
	if controllers.GetBalance(userA, models.USD) != 0 {
		t.Errorf("Transfer is not correct")
	}
	if controllers.GetBalance(userA, models.NGN) != 400 {
		t.Errorf("Transfer is not correct")
	}
	if controllers.GetBalance(userA, models.GBP) != 0.16 {
		t.Errorf("Transfer is not correct")
	}
}
