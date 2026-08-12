package main

import (
	"fmt"

	"mercuryops/internal/product"
)

func main() {
	item, err := product.New("Ноутбук", 450000, 0001)

	if err != nil {

		fmt.Println("Ошибка", err)
		return
	}

	fmt.Println("Товар:", item.Name())
	fmt.Println("Цена", item.Price())
	fmt.Println("SKU", item.SKU)
}
