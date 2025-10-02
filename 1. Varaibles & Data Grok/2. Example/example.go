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
		zeroValuesExample()
		typeConversionExample()
		constantsExample()
		multipleVaiables()
		commonerr()
}

/* 
	Breakdown:

	var customerName string = "David Beckham": Traditional declaration.
	orderID := "ORD123": Short declaration infers string.
	fmt.Printf: Uses %s (string), %.2f (float, 2 decimals), %t (bool).
	Output: Customer: David Beckham, Order: ORD123, Price: 150.75, Delivered: true
	- Why? Stores and displays order details for your app.
*/

//---------------------------------------

// Example 2: Zero Values

func zeroValuesExample() {
	var customerName string // Zero value: ""
	var orderID int // Zero value: 0
	var totalPrice float64 // Zero value: 0.0
	var isDelivered bool // Zero value: false
	
	fmt.Printf("Example 2: Zero Values:", customerName, orderID, totalPrice, isDelivered)
}

/*
	Breakdown:

	Variables declared without values get zero values.
	Prevents crashes if data isn’t set (e.g., no order price yet).
	Output: Zero values:  0 0 false
	- Why? Ensures your app handles unset data safely.
*/


//---------------------------------------


// Example 3: Type Conversion with Error Checking

func typeConversionExample() {
		priceStr := "89.99" // Text input from a website
    price, err := strconv.ParseFloat(priceStr, 64) // Convert to float64
    if err != nil { // Check if err is not nil (has an error)
        fmt.Println("Error converting price:", err)
        return // Stop program if conversion fails
    }
    fmt.Printf("Example3: Price: %.2f\n", price)
}

/* 
	Breakdown:

	import "strconv": Package for converting strings to numbers.
	price, err := strconv.ParseFloat(priceStr, 64): Converts "89.99" to float64.

	price: Gets 89.99.
	err: Gets error if conversion fails (e.g., "abc").


	if err != nil: Uses != (not equal) to check for errors.
	return: Stops if there’s an error.
	Output: Price: 89.99
	- Why? Converts user inputs for order calculations.
*/

//---------------------------------------

// Example 4: Constants and Naming

func constantsExample() {
	const DeliveryFee = 5.0
	orderPrice := 100.0
	total := orderPrice + DeliveryFee
	// Good morning
	deliveryAddress := "Tashken, Yunusabad"
	// Bad naming
	addr := "Tasheknt" // Unclear
	fmt.Printf("Example 4: Price: %.2f, Total: %.2f, Address: %.s\n", orderPrice, total, deliveryAddress)
	fmt.Println("Example 4: Bad name:", addr)
}

/* 
	Breakdown:

	const DeliveryFee: Fixed value.
	deliveryAddress: Clear name for readability.
	addr: Bad name (too vague).
	Output: Price: 100.00, Total: 105.00, Address: Tashkent, Yunusabad
	Bad name: Tashkent
	- Why? Shows good naming for maintainable code.	
*/

// ---------------------------------------

// Example 5: Multiple Variables and Comments

func multipleVaiables() {
	// Declare multiple varaibles at once
	var(
		customerName = "Ahmad Karimov"
		orderID = 67890
		totalPrice = 42.50
		isExpress = true
	)
	/* Print order details
		 for delivery tracking */
	fmt.Printf("Customer: %s, Order: %d, Price: %.2f, Express: %t", customerName, orderID, totalPrice, isExpress)
}

/* 
	Breakdown:

	var (...): Declares multiple variables.
	// and /*  Comments explain code.*/
	// Output: Customer: Ahmad Karimov, Order: 67890, Price: 42.50, Express: true
	// Why? Organizes order data efficiently.

// ---------------------------------------

// Example 6: Common Errors

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