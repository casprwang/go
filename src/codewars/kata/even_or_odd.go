package kata

func EvenOrOdd(number int) string {
	// your code here
	// Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
	if number&1 == 1 {
		return "Odd"
	} else {
		return "Even"
	}
}
