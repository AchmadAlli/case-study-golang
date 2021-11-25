package product

import "fmt"

type Product struct {
	name  string
	price int
}

var products = []Product{
	{
		name:  "Kasur",
		price: 1000000,
	},
	{
		name:  "Lemari",
		price: 500000,
	},
	{
		name:  "Meja",
		price: 300000,
	},
	{
		name:  "Kursi",
		price: 150000,
	},
	{
		name:  "Cermin",
		price: 100000,
	},
}

func ShowProducts() {
	fmt.Println("\nDaftar Produk")
	for i, product := range products {
		fmt.Printf("%d. %s \t\t %d\n", i+1, product.name, product.price)
	}
	fmt.Println("")
}
