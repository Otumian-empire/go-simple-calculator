package main

import (
	"fmt"
	"os"
)

// author: otumian-empire popecan1000@gmail.com
// project: simple calculator
// description: description and other details can be found in the README.md
// operation: get operands and operator (string) and return the results

// put negative second operands in brackets: (-23)
func blockOperand(operand int) string {
	if operand < 0 {
		return fmt.Sprintf("(%v)", operand)
	}

	return fmt.Sprintf("%v", operand)
}

// read and return operand
func getOperand(name string) int {
	fmt.Printf("Enter %v operand: ", name)
	var operand int
	fmt.Scan(&operand)

	fmt.Printf("The %v operand is %v\n\n", name, blockOperand(operand))

	return operand
}

// read operator
func getOperator() string {
	fmt.Print("Enter operator [+,-,*,/]: ")
	var operator string
	fmt.Scan(&operator)

	fmt.Printf("The operator is %v\n\n", operator)

	return operator
}

func exitCode(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func calculate(operator string, firstOperand, secondOperand int) int {
	var result int

	switch operator {
	case "+":
		result = firstOperand + secondOperand
	case "-":
		result = firstOperand - secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "/":
		if secondOperand == 0 {
			exitCode("Zero division for operator " + operator)
		}

		result = firstOperand / secondOperand
	default:
		exitCode("Operator not known")
	}

	return result
}

func displayResult(operator string, firstOperand, secondOperand, result int) {
	fmt.Printf("%v %v %v = %v", blockOperand(firstOperand), operator, blockOperand(secondOperand), result)
	fmt.Println("\n\nDone!!!")
}

func main() {
	fmt.Println("Simple Calculator")

	firstOperand := getOperand("first")
	secondOperand := getOperand("second")
	operator := getOperator()

	result := calculate(operator, firstOperand, secondOperand)

	displayResult(operator, firstOperand, secondOperand, result)
}
