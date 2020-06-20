package card

func TranslateMCC(code string) string {
	// представим, что mcc читается из файла (научимся позже)
	McResult := "Категория не найдена"
	MCC := map[string]string{
	//	"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5912": "Аптеки",
		"5812": "Кафе и рестораны",
	}
	for i := range MCC{
		if MCC[code] == MCC[i]{
			McResult = MCC[code]
		}
	}
	return McResult
}