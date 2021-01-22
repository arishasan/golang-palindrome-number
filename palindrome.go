package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {

	// Declare our variables
	var numFirst string
	var firstNumber, secondNumber, countedNumber string
	// End of declaration

	fmt.Println("----------------------------------------------")

	// Read and do validation for inputted number from user
	fmt.Println("Enter non-zero positive number (FirstNumber  (space) SecondNumber) : ")

	for {

		scanner := bufio.NewScanner(os.Stdin) // bufio can be use for long string and string with spaces in it
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		numFirst = scanner.Text() // Pass it to variable numFirst

		process := processPalindrome(numFirst)

		if process[0] == "err" {
			fmt.Println("Invalid inputted number")
		}else{

			if process[0] == "err-1" {
				fmt.Println("Inputted number can't be less than 1 and can't be more than 1.000.000.000")
			}else{

				if process[0] == "err-2" {
					fmt.Println("Inputted number can't be less than 1 and can't be more than 1.000.000.000. Also can't be less than first inputted number.")
				}else{
					firstNumber = process[0] // Primary Number
					secondNumber = process[1] // Secondary Number
					countedNumber = process[2] // Counted Palindrome Numbers
					break
				}

			}

		}

	}
	// End of input and validation


	fmt.Println("----------------------------------------------")
	fmt.Println("(INPUT) Start number :", firstNumber)
	fmt.Println("(INPUT) End number :", secondNumber)
	fmt.Println("(OUTPUT) Counted palindrome numbers :", countedNumber)	

}

// For UNIT TESTING only

func testingPalindrome(s string) int {

	process := processPalindrome(s)

	if process[0] == "err" {
		return 0
	}else{

		if process[0] == "err-1" {
			return 0
		}else{

			if process[0] == "err-2" {
				return 0
			}else{
				i, err := strconv.Atoi(process[2]) // Convert String to Integer
				if err != nil {}
				return i
			}

		}

	}

}

// End UNIT TESTING

func processPalindrome(s string) [3]string {

	var temporaryFirst, temporarySecond, countedPalindrome int
	var retArray [3]string

	words := strings.Split(s, " ") // Split the string into two with space delimeter between two numbers
	parseOne, err := strconv.Atoi(words[0]) // Take the first array from splitted string and convert it's type data to Integer
	parseTwo, err := strconv.Atoi(words[1]) // Take the second array from splitted string and convert it's type data to Integer

	if(err != nil){
		retArray[0] = "err"
		return retArray
	}

	if parseOne < 1 || parseOne > 1000000000 { // Condition if number less than 1 or more than 1 billion
		retArray[0] = "err-1"
		return retArray
	}else{

		if parseTwo < 1 || parseTwo > 1000000000 || parseTwo < parseOne { // Condition if number less than 1 or more than 1 billion and if secondary number is less than primary number
			retArray[0] = "err-2"
			return retArray
		}else{
			temporaryFirst = parseOne
			temporarySecond = parseTwo
		}
		
	}

	// Begin looping logic

	for i := temporaryFirst; i <= temporarySecond; i++ {

		if(isPalindrome(i)){ // If returned value is true add the value of countedPalindrome
			countedPalindrome++
		}

	}

	// End of looping logic

	// Assign to array, and then return it

	retArray[0] = strconv.Itoa(temporaryFirst) // Returned data as String
	retArray[1] = strconv.Itoa(temporarySecond) // Returned data as String
	retArray[2] = strconv.Itoa(countedPalindrome) // Returned data as String

	return retArray

	// End


}

// Logic function, determine whether the thrown number is palindrome number or not
func isPalindrome(x int) bool {

	if x < 0 || (x % 10 == 0 && x != 0) {
		return false // If x is 0 or null, return false already
	}

	var rev int
	for x > rev {
		rev = rev * 10 + x % 10
		x /= 10
	}

	return x == rev || x == rev / 10 // Consider it like odd and even cases.

}
// End of logic