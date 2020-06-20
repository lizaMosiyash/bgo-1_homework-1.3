package main

import (
	"fmt"
	card2 "github.com/lizaMosiyash/bgo-1_homework-1.3/card"
	card "github.com/lizaMosiyash/bgo-1_homework-1.3/pkg"
	"time"
)

func main() {
	master := &card.Card{
		Id: 1,
		Balance: 1_000_00,
		Currency: "RUB",
		Number: "1111_1111_1111_0000",
		Transactions: []*card.Transaction{
			&card.Transaction{
				Id:      1,
				Sum:     735_55,
				Date:    time.Date(2020, time.June, 9, 22, 06, 23, 0, time.UTC).Unix(),
				MccCode: "5411",
				Status:  "Операция в обработке",
			},
			&card.Transaction{
				Id:      2,
				Sum:     1_203_91,
				Date:    time.Date(2020, time.June, 9, 1, 46, 40, 0, time.UTC).Unix(),
				MccCode: "5812",
				Status:  "Операция в обработке",
			},
		},
	}
	fmt.Println(master)



	newTransaction := &card.Transaction{
		Id:      3,
		Sum:     999_99,
		Date:    time.Date(2020, time.June, 12, 1, 46, 40, 0, time.UTC).Unix(),
		MccCode: "5912",
		Status:  "Операция в обработке",
	}

	mccCodes := make([]string, 2, 2)
	mccCodes = append(mccCodes, "5812")
	mccCodes = append(mccCodes, "5912")
	mccCodes = append(mccCodes, "1111")
	card.AddTransaction(master, newTransaction)
	fmt.Println(card.SumByMCC(master.Transactions, mccCodes))

	category := card2.TranslateMCC(master.Transactions[0].MccCode)
	fmt.Println(category)

}
