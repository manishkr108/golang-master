package main

import (
	"fmt"
)

var Product = map[string]float64{
	"apple":  50,
	"banana": 30.50,
	"cocnut": 20.50,
}

var cart = make(map[string]int)

func addToCart(item string, qnt int) {
	if _, ok := Product[item]; !ok {
		fmt.Println("Product not found", item)
		return
	}

	if count, ok := cart[item]; ok {
		cart[item] = count + qnt
	} else {
		cart[item] = qnt
	}
	fmt.Printf("Added %d x %s to cart\n", qnt, item)
}

func removeCart(item string) {
	if _, ok := cart[item]; ok {
		delete(cart, item)
		fmt.Println("remove", item, "from cart")
	} else {
		fmt.Println("item not in cart", item)
	}
}

func viewCart() {
	if len(cart) == 0 {
		fmt.Println("you cart is empty")
		return
	}

	fmt.Println("you cart")

	total := 0.0

	for item, qnt := range cart {
		price := Product[item]
		subtotal := price * float64(qnt)

		fmt.Printf("- %s (x%d) = %.2f\n", item, qnt, subtotal)
		total += subtotal

	}
}

func main() {
	addToCart("apple", 5)

	addToCart("banana", 5)
	viewCart()
	addToCart("orange", 10)
	removeCart("banana")
	viewCart()
}
