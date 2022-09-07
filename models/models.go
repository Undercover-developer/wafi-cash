package models

const (
	USD  = "USD"
	GBP  = "GBP"
	NGN  = "NGN"
	YUAN = "YUAN"
)

type User struct {
	Id   int
	Name string
}
type Account struct {
	UserId int
	USD    float64
	GBP    float64
	NGN    float64
	YUAN   float64
}

type Transaction struct {
	FromUserId int
	ToUserId   int
	Amount     float64
	Currency   string
}
