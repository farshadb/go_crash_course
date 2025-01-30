package main

import (
	"fmt"
	"strconv"
	"strings"
)

// addThousandSeparator adds commas as thousand separators to a number
func addThousandSeparator[T ~int | ~int64 | ~float64](num T) string {
	switch v := any(num).(type) {
	case int:
		return formatInteger(v)
	case float64:
		return formatFloat(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// formatInteger formats an integer with thousand separators
func formatInteger(input interface{}) string {
	var s string
	switch v := input.(type) {
	case int:
		s = strconv.Itoa(v)
	case string:
		s = v
	default:
		return ""
	}

	// Initialize string builder with pre-allocated capacity
	var buf strings.Builder

	// The capacity is the length of the input string plus potential commas
	buf.Grow(len(s) + len(s)/3)

	// Track starting position for digit processing
	startOffset := 0

	// Handle negative numbers by writing minus sign
	// and adjusting start position
	if s[0] == '-' {
		buf.WriteByte('-')
		startOffset = 1
	}

	length := len(s)
	for i := startOffset; i < length; i++ {
		if (length-i)%3 == 0 && i != startOffset {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}

// formatFloat formats a float64 with thousand separators while preserving precision
func formatFloat(n float64) string {
	if n == 0 {
		return "0"
	}

	// Use strconv.FormatFloat to preserve original precision
	s := strconv.FormatFloat(n, 'f', -1, 64)
	parts := strings.Split(s, ".")

	// Format the integer part
	integerPart := formatInteger(mustAtoi(parts[0]))

	// Add decimal part if it exists
	if len(parts) > 1 {
		return fmt.Sprintf("%s.%s", integerPart, parts[1])
	}
	return integerPart
}

// mustAtoi converts string to int, panics on error (safe for our use case as we control the input)
func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err) // This should never happen given our input control
	}
	return n
}

// Person implements fmt.Stringer for better formatting
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

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

	fmt.Printf("Integer: %s\n", addThousandSeparator(1000000000000000000))   // Output: 1,000,000,000,000,000,000
	fmt.Printf("Float: %s\n", addThousandSeparator(1234567.89))              // Output: 1,234,567.89
	fmt.Printf("Large Float: %s\n", addThousandSeparator(1234567890000.891)) // Output: 1,234,567,890,000.891

	// Person struct formatting
	person := Person{"Farshad", 30}
	fmt.Printf("Person: %v\n", person) // Output: Farshad (30 years old)
}
