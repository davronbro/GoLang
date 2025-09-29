package main

/* 

📋 DAY 2 CHEATSHEET - Variables & Data Types

Variable Declaration Methods |

// Method 1: Traditional
		var customerName string
		var orderID int

// Method 2: Short (most popular)
		customerName := "Ahmad Karimov"
		orderID := 12345

// Method 3: Multiple at once
		var (
				name     = "Jane"
				orderID  = 678
				total    = 45.99
		)

==============================

Go Data Types for Delivery Systems |

// Strings - Text data
		customerName := "Bobur Umarov"
		address := "Tashkent, Chilanzar 12A"
		phone := "+998901234567"

// Integers - Whole numbers  
		orderID := 15847
		driverID := 203
		deliveryTime := 25  // minutes

// Float64 - Decimal numbers
		orderTotal := 89750.0    // UZS
		deliveryFee := 8000.0    // UZS
		distance := 7.2          // km
		rating := 4.8

// Booleans - True/False
		isDelivered := false
		isPaid := true
		isVIPCustomer := false
		isExpressDelivery := true

==================================

Zero Values (Default Values) |

var name string      // ""  (empty string)
var count int        // 0
var price float64    // 0.0  
var active bool      // false

==============================


Printf Formatting |

fmt.Printf("Order #%d\n", orderID)           // %d for integers
fmt.Printf("Customer: %s\n", customerName)   // %s for strings  
fmt.Printf("Total: %.2f UZS\n", total)       // %.2f for 2 decimal places
fmt.Printf("Express: %t\n", isExpress)       // %t for booleans

==============================

Variable Naming Rules |

// ✅ Good names
customerName := "Ahmad"
orderTotal := 45.99
deliveryAddress := "Tashkent"
isExpressDelivery := true

// ❌ Bad names
c := "Ahmad"        // Too short
x := 45.99          // Meaningless
addr := "Tashkent"  // Unclear
flag := true        // What flag?

==============================

Common Delivery Variables

// Customer data
		customerName := "Malika Karimova"
		phoneNumber := "+998901234567" 
		customerRating := 4.5

// Order data
		orderID := 15847
		restaurantID := 203
		orderTotal := 67500.0
		orderStatus := "preparing"

// Delivery data
		driverID := 45
		deliveryAddress := "Tashkent, Yunusabad 25B"
		estimatedTime := 30
		isExpressDelivery := true

// System data
		currentTime := "14:30"
		weatherCondition := "sunny"
		trafficLevel := "moderate"

*/