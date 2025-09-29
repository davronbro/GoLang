package main

import "fmt"

func main() {
    // Your code here
		fmt.Println("\n What's up everybody?")
		fmt.Print("Hello")           // Print without newline
		fmt.Println("Hello")         // Print with newline
		fmt.Printf("Hello") // Print with formatting
		company()
}

/* 
	Program Structure Rules

	1. Every file needs: package main (for executable programs)
	2. Every program needs: func main() (starting point)
	3. To print text: Import fmt and use fmt.Println()
	4. Strings: Always use double quotes "text"
	5. Comments: Use // for single line, -> /* */  /* -> for multiple lines


*/

func company() {
	// Company info
	fmt.Println("\n 🚀 FoodExpress Delivery System")

	// Service areas
	fmt.Println("📍 Service Areas: Downtown, University")

	// Operating hours
	fmt.Println("⏰ Open: 8:00 AM - 11:00 PM")

	// Contact info
	fmt.Println("📞 Call: 1-800-DELIVER")
}