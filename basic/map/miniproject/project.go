package main

import "fmt"

var products = map[string]float64{
	"apple":  30.50,
	"banana": 20,
	"milk":   22.50,
	"bread":  10,
}

var cart = make(map[string]int)

func addToCart(item string, qty int) {
	if _, ok := products[item]; !ok {
		fmt.Println("Product not found", item)
		return
	}
	if count, ok := cart[item]; ok {
		cart[item] = count + qty
	} else {
		cart[item] = qty
	}

	fmt.Printf("Added %d X %s to cart\n", qty, item)

}

func removeCart(item string) {
	if _, ok := cart[item]; ok {

		delete(cart, item)

		fmt.Println("Revome", item, "from cart")
	} else {
		fmt.Println("item not in cart", item)
	}
}

func viewCart() {
	if len(cart) == 0 {

		fmt.Println("cart is empty")
		return
	}

	fmt.Println("your cart")

	total := 0.0

	for item, qty := range cart {
		price := products[item]
		subtotal := price * float64(qty)
		fmt.Printf("- %s (x%d) = %.2f\n", item, qty, subtotal)
		total += subtotal

	}
	fmt.Printf("total %.2f\n\n", total)
}

func main() {
	addToCart("apple", 2)
	addToCart("banana", 5)
	addToCart("milk", 1)
	viewCart()

	removeCart("banana")
	viewCart()

	addToCart("chocolate", 1) // not in catalog
}
