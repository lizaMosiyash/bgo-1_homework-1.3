package card

type Transaction struct {
	Id      int64
	Sum     int64
	Date    int64
	MccCode string
	Status  string
}

type Card struct {
	Id           int
	Balance      int
	Currency     string
	Number       string
	Transactions []*Transaction
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

func SumByMCC(transactions []*Transaction, mcc []string) int64 {
	total := int64(0)
	for i := range transactions {
		for b := range mcc {
			if transactions[i].MccCode == mcc[b] {
				total += transactions[i].Sum
			}
		}
	}
	return total
}

func TranslateMCC(code string) string {
	// представим, что mcc читается из файла (научимся позже)
	McResult := "Категория не найдена"
	MCC := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5912": "Аптеки",
		"5812": "Кафе и рестораны",
	}
	value, ok := MCC[code]
	if ok == true {
		McResult = value
	}
	return McResult
}
