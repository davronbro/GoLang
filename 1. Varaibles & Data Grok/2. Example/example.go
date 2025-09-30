package main

// Example Code and Explanations

// Example 1: Declaring and Printing Variables

import (
	"fmt"
	"strconv"
)
func main() {
    // Traditional declaration
    var customerName string = "David Beckham"
    // Short declaration
    orderID := "ORD123"
    totalPrice := 150.75
    isDelivered := true
    fmt.Printf("Customer: %s, Order: %s, Price: %.2f, Delivered: %t\n", customerName, orderID, totalPrice, isDelivered)
    fmt.Println("Customer Name:", customerName)
    fmt.Println("Order ID:", orderID)
		fmt.Println("Total Price:", totalPrice)
    fmt.Println("Delivered:", isDelivered)
		constantsExample()
		typeConversionExample()
		combinedExample()
		commonerr()
}

/* 
	Breakdown:

	var customerName string = "David Beckham": Traditional declaration.
	orderID := "ORD123": Short declaration infers string.
	fmt.Printf: Uses %s (string), %.2f (float, 2 decimals), %t (bool).
	Output: Customer: David Beckham, Order: ORD123, Price: 150.75, Delivered: true
	Why? Stores and displays order details for your app.
*/

//---------------------------------------


// Example 2: Using Constants

func constantsExample() {
		const DeliveryFee = 5.0
    orderPrice := 100.0
    total := orderPrice + DeliveryFee
    fmt.Printf("Example 2: Price: %.2f, Delivery Fee: %.1f, Total with Fee: %.2f\n", orderPrice, DeliveryFee, total)
}

/* 
	Breakdown:

	const DeliveryFee = 5.0: Fixed fee, like JS const, can’t change.
	total := orderPrice + DeliveryFee: Adds fee to price.
	Output: Price: 100.00, Total with Fee: 105.00
	Why? Ensures consistent fees across orders, like in your company’s app.	
	
*/

//---------------------------------------

// Example 3: Type Conversion

func typeConversionExample() {
	priceStr := "89.99" // Input app/website
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Error", err)
		return // Stops program if conversion fails, like JS return
	}
	fmt.Printf("Example 3: Price: %.2f\n", price)
}

/* 
	Breakdown:

	priceStr := "89.99": Simulates text input (e.g., from a form).
	strconv.ParseFloat(priceStr, 64): Converts to float64 (64-bit precision), unlike JS parseFloat.
	err: Checks for invalid inputs (e.g., “abc”), with return to exit early.
	Output: Price: 89.99
	- Why? Converts user inputs to numbers for calculations, like payment totals in your app.	
*/

// ---------------------------------------

// Example 4: Combining Variables and Constants

func combinedExample() {
	const TaxRate = 0.1
	productionPrice := 200.0
	tax := productionPrice * TaxRate
	total := productionPrice + tax
	fmt.Printf("Example 4: Price: %.2f, Tax: %.2f, Total: %.2f\n", productionPrice, tax, total)
}

/* 
	Breakdown:

	const TaxRate = 0.1: Fixed 10% tax, like JS const.
	tax := productPrice * TaxRate: Calculates tax (200 * 0.1 = 20).
	total: Adds tax to price (200 + 20 = 220).
	Output: Price: 200.00, Tax: 20.00, Total: 220.00
  -	Why? Mimics calculating order totals with tax, common in fintech backends.
*/

// ---------------------------------------

// Example 5: Common Errors

func commonerr() {
    var price float64
    // price = "100.0" // Error: cannot use "100.0" (type string) as type float64
    price = 100.0
    fmt.Println(price)
    // const TaxRate = 0.1
    // TaxRate = 0.2 // Error: cannot assign to TaxRate
}

/* 
	Breakdown:

	price = "100.0": Causes type mismatch (like JS type coercion issues). Fix: Use strconv.ParseFloat.
	TaxRate = 0.2: Fails because constants are immutable. Fix: Use a variable.
	Output: 100
	- Why? Prevents bugs in backend code (e.g., mixing strings and numbers in order processing).
*/