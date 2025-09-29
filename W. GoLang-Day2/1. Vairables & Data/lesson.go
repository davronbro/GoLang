package main 

import "fmt"

// VARIABLES & DATA FOR DELIVERY SYSTEMS

/* 
	Why Variables Matter in Delivery Systems
	Yesterday you hardcoded everything. But real delivery apps need to handle changing data:

	 * Customer names change with every order
	 * Prices vary by restaurant and items
	 * Delivery addresses are always different
	 * Order status updates constantly

	Variables are like storage boxes that hold different types of information!

*/

// =================================================

// Variable Declaration in Go - Multiple Ways

	// Method 1: Traditional Declaration
	func main(){
		var globalCustomerName string
		var globalOrderID int
		var globalTotalPrice float64
		var globalIsDelivered bool
		
		fmt.Println("Zero values:", globalCustomerName, globalOrderID, 	globalTotalPrice, globalIsDelivered)

		// Call other functions
			short_dec()
			data_types()
			zero_values()
			naming()
	}

	// Method 2: Short Declaration (Most Popular)
	func short_dec() {

		customerName := "David Backham"        // string
		orderID := 12345                       // int  
		totalPrice := 25.99                    // float64
		isDelivered := false                   // bool

		fmt.Println("Short Declaration:", customerName, orderID, totalPrice, isDelivered)

	// Method 3: Multiple Variables at Once

		var (
			customerName2   = "David Backham"
			orderID2        = 67890
			totalPrice2     = 42.50
			isDelivered2    = true
		)

		fmt.Println("Multiple Varaibles at Once",customerName2, orderID2, totalPrice2, isDelivered2)
	}

// =================================================


// Go's Data Types for Delivery Systems
	func data_types () {

		// Strings - Text data
			customerName := "Ahmad Karimov"
			deliveryAddress := "Tashkent, Chilanzar District, 12A"
			phoneNumber := "+998901234567"
			specialInstructions := "Ring doorbell twice"

		// Integers - Whole numbers
			orderID := 15847
			restaurantID := 203
			driverID := 45
			deliveryTime := 25  // minutes

		// Float64 - Decimal numbers (prices, distances)
			orderTotal := 89.75
			deliveryFee := 4.99
			distance := 7.2     // kilometers
			driverRating := 4.8

		// Booleans - True/False values
			isDelivered := false
			isPaid := true
			isVIPCustomer := false
			isExpressDelivery := true

		fmt.Println(customerName, deliveryAddress, phoneNumber, specialInstructions, orderID, restaurantID, driverID, deliveryTime, orderTotal, deliveryFee, distance, driverRating, isDelivered, isPaid, isVIPCustomer, isExpressDelivery)
	}

// =================================================


//	Zero Values - Go's Safety Feature
		// When you declare variables without values, Go gives them "zero values":
		func zero_values() {
			var customerName string    // ""  (empty string)
			var orderID int           // 0
			var totalPrice float64    // 0.0
			var isDelivered bool      // false

			fmt.Println("Zero values:", customerName, orderID, totalPrice, isDelivered)

		// This prevents crashes - your delivery system won't break from uninitialized data!
		}

// =================================================

//	Variable Naming Best Practices
		func  naming(){
		// ✅ Good Delivery Variable Names:
			customerName := "Ahmad"
			orderTotal := 45.99
			deliveryAddress := "Tashkent, Yunusabad"
			isExpressDelivery := true

			fmt.Println(customerName, orderTotal, deliveryAddress, isExpressDelivery)

		// ❌ Bad Variable Names:
			c := "Ahmad"        // Too short
			x := 45.99          // Meaningless  
			addr := "Tashkent"  // Unclear abbreviation
			flag := true        // What kind of flag?

			fmt.Println(c, x, addr, flag)
		}