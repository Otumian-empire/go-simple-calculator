package main

import (
	"fmt"
)

// author: otumian popecan1000@gmail.com
// project: simple calculator
// description: description and other details can be found in the README.md
// operation: get operands and operator (string) and return the results

func main() {
	fmt.Println("Simple Calculator")

	// prompt user to enter operands
	fmt.Print("Enter first operand: ")
	var firstOperand int
	fmt.Scan(&firstOperand)

	fmt.Printf("The first operand is %v\n\n", firstOperand)

	// prompt user for second operand
	fmt.Print("Enter second operand: ")
	var secondOperand int
	fmt.Scan(&secondOperand)

	fmt.Printf("The second operand is %v\n\n", secondOperand)

	// prompt user for second operand
	fmt.Print("Enter second operand[+,-,*,/]: ")
	var operator string
	fmt.Scan(&operator)

	fmt.Printf("The operator is %v\n\n", operator)

	switch operator {
	case "+":
		fmt.Printf("%v %v %v = %v", firstOperand, operator, secondOperand, firstOperand+secondOperand)
	case "-":
		fmt.Printf("%v %v %v = %v", firstOperand, operator, secondOperand, firstOperand-secondOperand)
	case "*":
		fmt.Printf("%v %v %v = %v", firstOperand, operator, secondOperand, firstOperand*secondOperand)
	case "/":
		fmt.Printf("%v %v %v = %v", firstOperand, operator, secondOperand, firstOperand/secondOperand)
	default:
		fmt.Print("Operator not known")
	}

	fmt.Println("\n\nDone!!!")
}
