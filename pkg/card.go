package card

type Transaction struct {
	Id int64
	Sum int64
	Date int64
	MccCode string
	Status string
}

type Card struct {
	Id int
	Balance int
	Currency string
	Number string
	Transactions []*Transaction
}


func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

func SumByMCC(transactions []*Transaction, mcc []string) int64 {
	total := int64(0)
	for i := range transactions{
		for b := range mcc{
			if transactions[i].MccCode == mcc[b]{
				total += transactions[i].Sum
			}
		}
	}
	return total
}