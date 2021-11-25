package main

import (
	"fmt"
	"math"

	"github.com/AchmadAlli/case-study-golang/service"
)

const MENU_EXIT int = 0

func main() {
	const MENU_TRANSACTION int = 1
	const MENU_TRANSACTION_RESULT int = 2

	menu := math.MaxInt8

	for menu != MENU_EXIT {
		renderMenu()
		fmt.Print("Pilih Menu : ")
		_, err := fmt.Scanf("%d", &menu)

		if err != nil {
			fmt.Println("Masukkan angka sesuai menu lur ...")
			continue
		}

		switch menu {
		case MENU_TRANSACTION:
			handleTransaction()
		case MENU_TRANSACTION_RESULT:
			fmt.Println("ini hasilnya")
		default:
			continue
		}
	}
}

func renderMenu() {
	fmt.Println("Menu: ")
	fmt.Println("1. Transaksi ")
	fmt.Println("2. Laporan Transaksi ")
	fmt.Println("0. Keluar ")
}

func handleTransaction() {
	var productId = math.MaxInt8
	var qty int

	for productId != MENU_EXIT {
		service.ShowProducts()
		fmt.Print("Pilih Produk : ")
		_, err := fmt.Scanf("%d", &productId)

		if err != nil {
			fmt.Println("Wrong Input")
		}

		if productId == 0 {
			continue
		}

		fmt.Print("Pilih Produk : ")
		_, _ = fmt.Scanf("%d", &qty)

		product := service.GetProduct(productId)
		service.AddToCart(product, qty)
	}

	service.ShowCarts()
}
