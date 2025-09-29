package main

import "fmt"

func main(){
	// We'll build this step by step

// Step 1: Customer Data
customerName := "Davron"
phoneNumber := "+998337070099"
customerType := "regular" // regular, vip, new
customerRating := 4.5

// ======================

// Step 2: Order Details
orderID := 15847
restaurantName := "Besh Qozon"
restaurantRating := 4.7
orderItems := "Full set choyxona osh, salat, non"
baseAmount := 65000.0 // UZS

// ======================

// Step 3: Delivery Calculation System
deliveryAddress := "Tashkent, Uch tepa District G9a 4-12"
distance := 5.5 // kilometers
weatherCondition := "sunny" // sunny, rainy, snowy
trafficLevel := "moderate" // light, moderate, heavy
isRushhour := true
isWeekend := false

// ======================

// Step 5: Smart Fee Calculator

// Base delivery fee calculation
var baseFee float64

if distance <= 3.0 {
	baseFee = 19000.0
}else if distance <= 7.0{
	baseFee = 29000.0
}else{
	baseFee = 31000.0
}
// Apply multipliers
var finalDeliveryFee = baseFee

if weatherCondition == "rainy"{
	finalDeliveryFee = finalDeliveryFee * 1.2
}else if weatherCondition == "snowy"{
	finalDeliveryFee = finalDeliveryFee * 1.5
}else{
	finalDeliveryFee = finalDeliveryFee * 1
}

// Traffic Conditions
if trafficLevel == "heavy"{
	finalDeliveryFee = finalDeliveryFee * 1.2
}

// Rush Hour Surcharge
if isRushhour{
	finalDeliveryFee = finalDeliveryFee * 1.4
}

// VIP discount

if customerType == "vip"{
	finalDeliveryFee = finalDeliveryFee * 0.7
}

// ======================

// Step 6. Final Calculations and Display

// Calculation Totals

finalTotal := baseAmount + finalDeliveryFee
estimatedeliveryTime := int(distance * 3) // 3 minutes per km base

// Adjust time and conditions

if trafficLevel == "heavy" {
	estimatedeliveryTime = estimatedeliveryTime * 2
}

if weatherCondition == "rainy"{
	estimatedeliveryTime = estimatedeliveryTime + 10
}

	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║        ORDER CONFIRMATION            ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Printf("Order ID: %d\n", orderID)
	fmt.Printf("Customer: %s (%s) | Rating: %.1f\n", customerName, customerType, customerRating)
	fmt.Printf("Phone: %s\n", phoneNumber)
	fmt.Printf("Restaturant: %s (%.1f)\n", restaurantName, restaurantRating)
	fmt.Printf("Itmes: %s\n", orderItems)
	fmt.Printf("Delivery Address: %s\n", deliveryAddress)
	fmt.Printf("Weather: %s | Traffic: %s\n", weatherCondition, trafficLevel)
	fmt.Printf("Rush Hour: %t | Weekend: %t\n", isRushhour, isWeekend)
	fmt.Println("────────────────────────────────────────")
	fmt.Printf("Base Amount: %.1f UZS\n", baseAmount)
	fmt.Printf("Delivery Fee: %.1f UZS\n", finalDeliveryFee)
	fmt.Printf("TOTAL: %.1f UZS\n", finalTotal)
	fmt.Printf("Estimated Delivery: %d minutes\n", estimatedeliveryTime)
}







