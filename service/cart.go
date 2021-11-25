package service

import "fmt"

type Cart struct {
	product    Product
	quantity   int
	totalPrice int
}

var carts = []Cart{}

func AddToCart(p Product, qty int) {
	total := p.price * qty
	carts = append(carts, Cart{p, qty, total})
	fmt.Print("Produk telah disimpan\n")
}

func Checkout() []Cart {
	var total int
	fmt.Printf("\nberikut daftar produk yang dibeli\n")
	fmt.Printf("%s. %s \t %s \t\t %s \t %s \n", "no", "produk", "harga", "jumlah", "total harga")
	for i, item := range carts {
		fmt.Printf("%d. %s \t %d \t %d \t\t %d \n", i+1, item.product.name, item.product.price, item.quantity, item.totalPrice)
		total += item.totalPrice
	}

	fmt.Printf("\n\nTotal biaya yang harus dibayar adalah : Rp %d \n", total)
	result := carts
	carts = []Cart{}
	return result
}

func GetTotalPrice(carts []Cart) int {
	var total int
	for _, item := range carts {
		total += item.product.price * item.quantity
	}

	return total
}
