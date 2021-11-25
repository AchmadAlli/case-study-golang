package main

import (
	"fmt"
	"math"

	product "github.com/AchmadAlli/case-study-golang/modules"
)

func main() {
	const MENU_TRANSACTION int = 1
	const MENU_TRANSACTION_RESULT int = 2
	const MENU_EXIT int = 0

	menu := math.MaxInt8

	for menu != MENU_EXIT {
		renderMenu()
		_, err := fmt.Scanf("%d", &menu)

		if err != nil {
			fmt.Println("Masukkan angka sesuai menu lur ...")
			continue
		}

		switch menu {
		case MENU_TRANSACTION:
			product.ShowProducts()
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
