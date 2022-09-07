package test

import (
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
