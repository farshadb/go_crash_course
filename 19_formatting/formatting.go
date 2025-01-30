package main

import "fmt"

func main() {
	// Basic Formatting
	fmt.Printf("%%v (default format): %v\n", 42)          // Output: 42
	fmt.Printf("%%+v (with field names): %+v\n", struct { // Output: {Field:42}
		Field int
	}{42})
	fmt.Printf("%%#v (Go syntax): %#v\n", 42)       // Output: 42
	fmt.Printf("%%T (type): %T\n", 42)              // Output: int
	fmt.Printf("%%%% (literal percent sign): %%\n") // Output: %
	fmt.Printf("%%t (boolean): %t\n", true)         // Output: true

	// Integer Formatting
	fmt.Printf("%%b (binary): %b\n", 42)            // Output: 101010
	fmt.Printf("%%c (Unicode character): %c\n", 65) // Output: A
	fmt.Printf("%%d (decimal): %d\n", 42)           // Output: 42
	fmt.Printf("%%o (octal): %o\n", 42)             // Output: 52
	fmt.Printf("%%x (hex lowercase): %x\n", 42)     // Output: 2a
	fmt.Printf("%%X (hex uppercase): %X\n", 42)     // Output: 2A
	fmt.Printf("%%U (Unicode format): %U\n", 65)    // Output: U+0041

	// Floating-Point Formatting
	fmt.Printf("%%f (decimal): %f\n", 3.14159)                // Output: 3.141590
	fmt.Printf("%%.2f (2 decimal places): %.2f\n", 3.14159)   // Output: 3.14
	fmt.Printf("%%e (scientific notation): %e\n", 123.456)    // Output: 1.234560e+02
	fmt.Printf("%%E (scientific uppercase): %E\n", 123.456)   // Output: 1.234560E+02
	fmt.Printf("%%g (compact representation): %g\n", 123.456) // Output: 123.456
	fmt.Printf("%%G (compact uppercase): %G\n", 123.456)      // Output: 123.456

	// String and Byte Formatting
	fmt.Printf("%%s (plain string): %s\n", "hello")  // Output: hello
	fmt.Printf("%%q (quoted string): %q\n", "hello") // Output: "hello"
	fmt.Printf("%%x (hex lowercase): %x\n", "hello") // Output: 68656c6c6f
	fmt.Printf("%%X (hex uppercase): %X\n", "hello") // Output: 68656C6C6F

	// Pointer Formatting
	x := 42
	fmt.Printf("%%p (pointer address): %p\n", &x) // Output: 0xc0000... (memory address)

	// Width and Precision
	fmt.Printf("%%5d (width 5): %5d\n", 42)                            // Output:    42
	fmt.Printf("%%05d (width 5, zero-padded): %05d\n", 42)             // Output: 00042
	fmt.Printf("%%8.2f (width 8, 2 decimal places): %8.2f\n", 3.14159) // Output:     3.14
	fmt.Printf("%%.3s (max 3 characters): %.3s\n", "hello")            // Output: hel

	// Special Characters
	fmt.Printf("Newline: \\n\n")         // Output: Newline: (newline)
	fmt.Printf("Tab: \\t\n")             // Output: Tab:    (tab)
	fmt.Printf("Backslash: \\\\\n")      // Output: Backslash: \
	fmt.Printf("Double quote: \\\"\n")   // Output: Double quote: "
	fmt.Printf("Carriage return: \\r\n") // Output: Carriage return: (carriage return)
	fmt.Printf("Backspace: \\b\n")       // Output: Backspace: (backspace)
	fmt.Printf("Form feed: \\f\n")       // Output: Form feed: (form feed)
	fmt.Printf("Alert (bell): \\a\n")    // Output: Alert (bell): (bell sound)
	fmt.Printf("Vertical tab: \\v\n")    // Output: Vertical tab: (vertical tab)

	// Advanced Formatting
	fmt.Printf("%%-10s (left-justified): %-10s|\n", "hello") // Output: hello     |
	fmt.Printf("%%+d (explicit sign): %+d\n", 42)            // Output: +42
	fmt.Printf("%% d (space padding): % d\n", 42)            // Output:  42
	fmt.Printf("%%'d (digit grouping): %'d\n", 1000000)      // Output: 1,000,000

	// Custom Formatting with fmt.Stringer
	type Person struct {
		Name string
		Age  int
	}

	person := Person{"Farshad", 30}
	fmt.Printf("%%v (custom String()): %v\n", person) // Output: Farshad (30 years old)
}
