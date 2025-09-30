package main

// What Are Variables & Data Types?
/* 
 1. Variables:  
		Variables: Containers for storing data, like a customer’s name or order price in your delivery app. 
		You declare them with var (e.g., var price float64 = 100.0) or := (e.g., price := 100.0) inside a function like main().


 2. Data Types: Define data kind: 
		string: Text (e.g., customer name like "David Beckham" or order ID like "ORD123").
		int: Whole numbers (e.g., order quantity like 5).
		float64: Decimals (e.g., order total like 99.99).
		bool: True/false (e.g., is order delivered?).


 3. Constants: Fixed values that can’t change, like const DeliveryFee = 5.0. Used for unchanging rules (e.g., fees or tax rates).


 4. Zero Values: When you declare a variable without a value (e.g., var price float64), Go assigns a default “zero value”:
		string: "" (empty string).
		int: 0.
		float64: 0.0.
		bool: false.
		This prevents crashes in your delivery app if data isn’t set yet.


 5. Type Conversion: Changing types, like turning a string ("89.99") into a float64 for calculations, using the strconv package.


 6. Error Checking: 
 		Using an error variable (err) to catch problems (e.g., invalid inputs). 
 		nil means no error; != (not equal) checks if there’s an error.


 7. Program Structure:
		Every Go file starts with package main (for programs that run).
		The program starts at func main().
		Use import to include packages like "fmt" (for printing) or "strconv" (for conversions). 


 8. Comments: Notes in code to explain what it does.
	
		// Single-line: // This is a comment.
	 	// Multi-line: /* This is a multi-line comment */

/*
		-Why This Matters: Variables store order data (e.g., prices, customer names), 
		constants ensure consistent fees, zero values prevent errors, 
		and type conversion handles user inputs (e.g., prices from a website). 
		These are key for your delivery app’s backend to process orders reliably.
*/

// ------------------------------------------------------------------------

// Key Operations

/* 
	1. Declaring Variables:
		Traditional: var customerName string = "David Beckham".
		Short: customerName := "David Beckham" (infers string, only inside func).
		Multiple at once: var ( name = "Ahmad"; price = 99.99 ).


	2. Constants: const TaxRate = 0.1 (can’t change).


	3. Type Conversion: strconv.ParseFloat("89.99", 64) converts string to float64.
		strconv: Package for conversions.
		Returns price (number) and err (error if conversion fails).



	4. Error Checking:
		err: Stores error messages.
		nil: Means no error.
		!=: Checks if err is not nil (has an error).
		return: Stops the program if there’s an error.

 
	5. Printing:
		fmt.Println(x): Prints with a newline.
		fmt.Printf("Price: %.2f\n", x): Formats output (e.g., %.2f for 2 decimals).

	
	6. Variable Naming:
		Good: Clear names like customerName, orderTotal, isExpressDelivery.
		Bad: Vague names like c, x, flag.
*/

// ---------------------------------------



