// Real Examples

// Example 1: Basic Order Data Storage

package main

import "fmt"

func main() {
    // Customer information
    customerName := "Bobur Umarov"
    phoneNumber := "+998901234567"
    
    // Order details
    orderID := 15847
    restaurantName := "Plov House Tashkent"
    orderTotal := 45000.0  // UZS
    
    // Delivery information
    deliveryAddress := "Tashkent, Sergeli District, Building 15A"
    isExpressDelivery := true
    
    // Display order summary
    fmt.Println("\nExample 1: Basic Order Data Storage\n " )
    fmt.Println("=== ORDER CONFIRMATION ===")
    fmt.Printf("Order ID: #%d\n", orderID)
    fmt.Printf("Customer: %s\n", customerName)
    fmt.Printf("Phone: %s\n", phoneNumber)
    fmt.Printf("Restaurant: %s\n", restaurantName)
    fmt.Printf("Total: %.1f UZS\n", orderTotal)
    fmt.Printf("Delivery Address: %s\n", deliveryAddress)
    fmt.Printf("Express Delivery: %t\n", isExpressDelivery)
    fmt.Printf("======================\n")
    orders()
    calculation()
}


// ===========================

// Example 2: Multiple Orders Processing


func orders (){
    // Order1
    order1ID := 1001
    order1Customer := "Malika Karimova"
    order1Total := 67500.0

    //Order2
    order2ID := 1002
    order2Customer := "Sardor Rahimov"
    order2Total := 23000.0

    //Order3
    order3ID := 1003
    order3Customer := "Nilufar Yusupova"
    order3Total := 89000.0

    // Calculate Daily Statistics
    totalOrders := 3
    dailyRevenue := order1Total + order2Total + order3Total
    averageOrderValue := dailyRevenue / float64 (totalOrders)

    fmt.Println("\nExample 2: Multiple Order Processing\n " )
    fmt.Println("=== DAILY DELIVERY SUMMARY ===")
    fmt.Printf("Total Orders: %d\n", totalOrders)
    fmt.Printf("Daily Revenue: %.1f UZS\n", dailyRevenue)
    fmt.Printf("Average Order Value: %.1f UZS\n", averageOrderValue)

    fmt.Println("\n===ORDER DETAILS===")
    fmt.Printf("Order #%d - %s: %.1f UZS\n", order1ID, order1Customer, order1Total)
    fmt.Printf("Order #%d - %s: %.1f UZS\n", order2ID, order2Customer, order2Total)
    fmt.Printf("Order #%d - %s: %.1f UZS\n", order3ID, order3Customer, order3Total)
 }  


 // ===========================

 // Example 3: Delivery Fee Calculation
 
 func calculation(){
    // Customer order
    customerName := "Jasur Abdullayev"
    baseOrderAmount := 50000.0 // UZS

    // Delivery paprameters
    distance := 8.5 // kilometers
    isRushHour := true // if it's not Rush Hour you do it like "false"
    isWeekend := false // if it's Weekend you can do it like "ture"

    // Calculate Delivery fee
    var deliveryFee float64

    // Base Delivery fee
    if distance <= 3.0 {
        deliveryFee = 19000.0
    }else if distance <= 7.0 {
        deliveryFee = 28250.0      
    }else{
        deliveryFee = 31000.0
    }

    // Rush hour surcharge
    if isRushHour{
        deliveryFee = deliveryFee * 1.5
    }

    if isWeekend{
        deliveryFee = deliveryFee * 0.8
    }

    // Calculate final total
    finalTotal := baseOrderAmount + deliveryFee
        
    fmt.Println("\nExample 3: Delivery Fee Calculation\n " )
    fmt.Println("===DELIVERY FEE CALCULATION===")
    fmt.Printf("Customer: %s\n", customerName)
    fmt.Printf("Base Order: %.1f UZS\n", baseOrderAmount)
    fmt.Printf("Distance: %.1f km\n", distance)
    fmt.Printf("Rush Hour: %t\n", isRushHour)
    fmt.Printf("Weekend: %t\n", isWeekend)
    fmt.Printf("Delivery Fee: %.1f UZS\n", deliveryFee)
    fmt.Printf("Final Total: %.1f UZS\n", finalTotal)
 }
