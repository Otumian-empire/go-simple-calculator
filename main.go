package main

import (
	"fmt"
	"os"
	"strings"
)

// author: otumian-empire popecan1000@gmail.com
// project: simple calculator
// description: description and other details can be found in the README.md
// operation: get operands and operator (string) and return the results

// a func which takes in two int args and returns an int
type BinaryFuncInt func(int, int) int

const (
	OPERATORS = "+,-,*,/"
	ADD       = "+"
	DIV       = "/"
	MUL       = "*"
	SUB       = "-"
)

const ZERO = 0

// put negative second operands in brackets: (-23)
func blockOperand(operand int) string {
	if operand < ZERO {
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
	fmt.Printf("Enter operator [%v]: ", OPERATORS)
	var operator string
	fmt.Scan(&operator)

	fmt.Printf("The operator is %v\n\n", operator)

	return operator
}

func exitCode(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func displayResult(operator string, firstOperand, secondOperand, result int) {
	fmt.Printf("%v %v %v = %v\n\n", blockOperand(firstOperand), operator, blockOperand(secondOperand), result)
	fmt.Println("Done!!!")
}

func add(x, y int) int { return x + y }
func div(x, y int) int { return x / y }
func mul(x, y int) int { return x * y }
func sub(x, y int) int { return x - y }

func calculate() map[string]BinaryFuncInt {
	return map[string]BinaryFuncInt{
		ADD: add,
		SUB: sub,
		MUL: mul,
		DIV: div,
	}
}

func validateOperator(op string) bool {
	return strings.Contains(OPERATORS, op)
}

func main() {
	fmt.Println("Simple Calculator")

	firstOperand := getOperand("first")
	secondOperand := getOperand("second")
	operator := getOperator()

	// validate operator
	if !validateOperator(operator) {
		exitCode("Operator not known")
	}

	// if the operator is / then check for zero division
	if secondOperand == ZERO {
		exitCode("Zero division for operator " + operator)
	}

	// compute the result
	result := calculate()[operator](firstOperand, secondOperand)

	displayResult(operator, firstOperand, secondOperand, result)
}
