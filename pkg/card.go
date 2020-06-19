package card

type Transaction struct {
	Id int64
	Sum int
	Date int64
	MccCode string
	Status string
}

type Card struct {
	Id int
	Issuer string
	Balance int
	Currency string
	Number string
	Transactions []*Transaction
}

//func AddTransaction(card card.Transactions, transaction *Transaction) {
//	card = append()
//}

//func SumByMCC(transactions []Transaction, mcc []string) int64 {
	// TODO: ваш код
//}