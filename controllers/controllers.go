package controllers

import (
	"math/rand"
	"reflect"
	"wafi-cash/models"
)

var users []models.User
var accounts []models.Account

var usdExchangeRate = map[string]float64{"USD": 1, "NGN": 415, "GBP": 0.86, "YUAN": 6.89}

func AddUser(name string) models.User {
	id := rand.Intn(100)
	newUser := models.User{
		Id:   id,
		Name: name,
	}
	AddAccount(newUser)
	users = append(users, newUser)
	return newUser
}

func AddAccount(user models.User) {
	accounts = append(accounts, models.Account{
		UserId: user.Id,
		USD:    0,
		GBP:    0,
		NGN:    0,
		YUAN:   0,
	})
}

func GetAccounts() []models.Account {
	return accounts
}

func Deposit(user models.User, amount float64, currency string) {
	for i, account := range accounts {
		if account.UserId == user.Id {
			switch currency {
			case models.USD:
				accounts[i].USD += amount
			case models.GBP:
				accounts[i].GBP += amount
			case models.NGN:
				accounts[i].NGN += amount
			case models.YUAN:
				accounts[i].YUAN += amount
			}
		}
	}
}

func GetBalance(user models.User, currency string) float64 {
	for _, account := range accounts {
		if account.UserId == user.Id {
			switch currency {
			case models.USD:
				return account.USD
			case models.GBP:
				return account.GBP
			case models.NGN:
				return account.NGN
			case models.YUAN:
				return account.YUAN
			}
		}
	}
	return 0
}

func Transfer(fromUserId int, toUserId int, amount float64, currency string) {
	for i, account := range accounts {
		if account.UserId == fromUserId {
			switch currency {
			case models.USD:
				if account.USD >= amount {
					accounts[i].USD -= amount
				} else {
					performAggregateTransaction(i, amount, currency, usdExchangeRate)
				}
			case models.GBP:
				if account.GBP >= amount {
					accounts[i].GBP -= amount
				} else {
					//performAggregateTransaction with GBP exchange rate
					return
				}
			case models.NGN:
				if account.NGN >= amount {
					accounts[i].NGN -= amount
				} else {
					//performAggregateTransaction with NGN exchange rate
					return
				}
			case models.YUAN:
				if account.YUAN >= amount {
					accounts[i].YUAN -= amount
				} else {
					//performAggregateTransaction with YUAN exchange rate
					return
				}
			}
		}
		if account.UserId == toUserId {
			switch currency {
			case models.USD:
				accounts[i].USD += amount
			case models.GBP:
				accounts[i].GBP += amount
			case models.NGN:
				accounts[i].NGN += amount
			case models.YUAN:
				accounts[i].YUAN += amount
			}
		}
	}
}

func isPossibleTransaction(account models.Account, amount float64, currency string, currencyExchangeRate map[string]float64) bool {
	balance := reflect.ValueOf(account)
	for i := 0; i < balance.NumField(); i++ {
		if balance.Type().Field(i).Name == "UserId" {
			continue
		}
		if balance.Type().Field(i).Name == currency {
			balance.Field(i).SetFloat(balance.Field(i).Float() - amount)
			amount = amount - balance.Field(i).Float()
		} else {
			balance.Field(i).SetFloat((balance.Field(i).Float() * currencyExchangeRate[currency]) - amount)
			amount = amount - (balance.Field(i).Float() * currencyExchangeRate[currency])
		}
		if amount <= 0 {
			return true
		}
	}
	return false
}

func performAggregateTransaction(accountIdx int, amount float64, currency string, currencyExchangeRate map[string]float64) {
	if isPossibleTransaction(accounts[accountIdx], amount, currency, currencyExchangeRate) {
		balance := reflect.ValueOf(accounts[accountIdx])
		for i := 0; i < balance.NumField(); i++ {
			if balance.Type().Field(i).Name == "UserId" {
				continue
			}

			if balance.Type().Field(i).Name == currency {
				balance.Field(i).SetFloat(balance.Field(i).Float() - amount)
				amount = amount - balance.Field(i).Float()
			} else {
				balance.Field(i).SetFloat((balance.Field(i).Float() * currencyExchangeRate[currency]) - amount)
				amount = amount - (balance.Field(i).Float() * currencyExchangeRate[currency])
			}
			if amount <= 0 {
				return
			}
		}
	} else {
		return
	}
}
