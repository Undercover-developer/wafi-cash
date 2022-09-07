package main

/*
	All the code was put in the same file for simplicity and easy inspection.
*/
import (
	"context"
	"fmt"
	"time"
	"wafi-cash/controllers"
	"wafi-cash/models"
)

func main() {
	trxCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Add usersA
	userA := controllers.AddUser("User A")
	//deposit 10 dollars
	controllers.Deposit(userA, 10, models.USD)
	// Add usersB
	userB := controllers.AddUser("User B")
	//deposit 20 dollars
	controllers.Deposit(userB, 20, models.USD)
	// User B sends 15 dollars to User A
	trxChannel := make(chan models.Transaction) // to simulate transaction queue
	trx := models.Transaction{
		FromUserId: userB.Id,
		ToUserId:   userA.Id,
		Amount:     15,
		Currency:   models.USD,
	}
	//schedule transaction queue handler
	go func() {
		for {
			select {
			case <-trxCtx.Done():
				return
			case trx := <-trxChannel:
				controllers.Transfer(trx.FromUserId, trx.ToUserId, trx.Amount, trx.Currency)
			}
		}
	}()
	trxChannel <- trx

	// User A checks their balance and has 25 dollars
	fmt.Println("User A balance is: ", controllers.GetBalance(userA, models.USD))
	// User B checks their balance and has 5 dollars
	fmt.Println("User B balance is: ", controllers.GetBalance(userB, models.USD))
	// User A transfers 25 dollars from their account
	trx = models.Transaction{
		FromUserId: userA.Id,
		ToUserId:   0,
		Amount:     25,
		Currency:   models.USD,
	}
	trxChannel <- trx
	time.Sleep(time.Second)
	// User A checks their balance and has 0 dollars
	fmt.Println("User A balance is: ", controllers.GetBalance(userA, models.USD))
}
