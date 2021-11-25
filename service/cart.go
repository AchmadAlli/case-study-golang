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

func showCarts() {

}
