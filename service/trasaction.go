package service

type Transaction struct {
	customer string
	detail   []Cart
	total    int
}

var transactions = []Transaction{}

func AddTransaction(customer string, carts []Cart) {
	transactions = append(transactions, Transaction{
		customer: customer,
		detail:   carts,
		total:    GetTotalPrice(carts),
	})
}
