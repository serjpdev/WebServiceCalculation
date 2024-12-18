package calculation

import (
	"errors"
	"regexp"
	"strconv"
)

func checkOperation(symbol string) bool {
	switch symbol {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	default:
		return false
	}
}

func getWeightOp(operator string) int {
	switch operator {
	case "(":
		return 0
	case ")":
		return 1
	case "+":
		return 2
	case "-":
		return 2
	case "*":
		return 3
	case "/":
		return 3
	default:
		return -1
	}
}

func calculateInfix(operator string, operand1, operand2 float64) float64 {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		return 0
	}
}

func getNumberFromString(expression string, pos *int) string {
	var output string

	for ; *pos < len(expression); *pos++ {
		symbol := string((expression)[*pos])
		_, err := strconv.Atoi(symbol)

		if err == nil || symbol == "." {

			output += symbol
		} else {

			*pos--
			break
		}
	}

	return output
}

func convertInfixToPostfix(infixExpression string) (string, error) {
	var postfixExpression string // Выходная строка
	var operationsStack []string // Стек операторов (так-то массив, но будет вести себя как стек)

	if _, err := validateInfix(infixExpression); err != nil {

		return postfixExpression, err
	}

	for i := 0; i < len(infixExpression); i++ {
		symbol := string((infixExpression)[i])

		if _, err := strconv.Atoi(symbol); err == nil || symbol == "." {

			postfixExpression += getNumberFromString(infixExpression, &i) + " "
		} else if checkOperation(symbol) {

			if len(operationsStack) == 0 ||
				getWeightOp(operationsStack[len(operationsStack)-1]) < getWeightOp(symbol) {

				operationsStack = append(operationsStack, symbol)
			} else {

				for len(operationsStack) > 0 {
					if getWeightOp(operationsStack[len(operationsStack)-1]) < getWeightOp(symbol) {

						break
					}

					postfixExpression += operationsStack[len(operationsStack)-1] + " "

					operationsStack = operationsStack[:len(operationsStack)-1]
				}

				if len(operationsStack) == 0 ||
					getWeightOp(operationsStack[len(operationsStack)-1]) < getWeightOp(symbol) {
					operationsStack = append(operationsStack, symbol)
				}
			}
		} else if symbol == "(" {

			operationsStack = append(operationsStack, symbol)
		} else if symbol == ")" {

			for len(operationsStack) > 0 {
				if operationsStack[len(operationsStack)-1] == "(" {

					break
				}

				postfixExpression += operationsStack[len(operationsStack)-1] + " "

				operationsStack = operationsStack[:len(operationsStack)-1]
			}

			operationsStack = operationsStack[:len(operationsStack)-1]

		}
	}

	for len(operationsStack) > 0 {
		postfixExpression += operationsStack[len(operationsStack)-1] + " "
		operationsStack = operationsStack[:len(operationsStack)-1]
	}

	return postfixExpression, nil
}

func checkParentheses(expression string) bool {
	var parenthesesStack []string

	for i := 0; i < len(expression); i++ {
		symbol := string((expression)[i])

		if symbol == "(" {
			parenthesesStack = append(parenthesesStack, symbol)
			continue
		}
		if symbol != ")" {
			continue
		}

		if len(parenthesesStack) == 0 {

			return false
		}

		parenthesesStack = parenthesesStack[:len(parenthesesStack)-1]
	}

	return len(parenthesesStack) == 0
}

func checkBinaryOperations(expression string) bool {
	operandsCount, operatorsCount := 0, 0

	for i := 0; i < len(expression); i++ {
		symbol := string((expression)[i])
		_, err := strconv.Atoi(symbol)

		if err == nil || symbol == "." {
			getNumberFromString(expression, &i)
			operandsCount++
		} else if checkOperation(symbol) {
			operatorsCount++
		}
	}

	return operandsCount-operatorsCount == 1
}

func checkSymbols(expression string) bool {
	// spaces should be in start of pattern
	pattern := `^[0-9\.\+\-\*\/()\s]+$`
	return regexp.MustCompile(pattern).MatchString(expression)
}

func validateInfix(expression string) (bool, error) {
	// validation infix expression
	if len(expression) == 0 {

		return false, errors.New("expression is empty")
	}
	if !checkSymbols(expression) {

		return false, errors.New("expression contains invalid characters")
	}
	if !checkParentheses(expression) {

		return false, errors.New("parentheses is not valid")
	}
	if !checkBinaryOperations(expression) {
		return false, errors.New("binary operations is not valid")
	}

	return true, nil
}

func Calc(expression string) (float64, error) {
	postfixExpression, err := convertInfixToPostfix(expression)

	if err != nil {
		return 0, err
	}
	// stack of  operands
	var operandsStack []float64

	// read postfix entry by symbol
	for i := 0; i < len(postfixExpression); i++ {
		symbol := string(postfixExpression[i])

		if _, err := strconv.Atoi(symbol); err == nil || symbol == "." {
			// check if symbol is digit or dot,
			// add to stack

			operand, err := strconv.ParseFloat(getNumberFromString(postfixExpression, &i), 64)
			if err != nil {
				return 0, err
			}

			operandsStack = append(operandsStack, operand)
		} else if checkOperation(symbol) {
			// if symbol is operator

			var operand1, operand2 float64

			// pop 2 last element
			if len(operandsStack) > 0 {
				operand2 = operandsStack[len(operandsStack)-1]
				operandsStack = operandsStack[:len(operandsStack)-1]
			}
			if len(operandsStack) > 0 {
				operand1 = operandsStack[len(operandsStack)-1]
				operandsStack = operandsStack[:len(operandsStack)-1]
			}

			if symbol == "/" && operand2 == 0.0 {
				return 0, errors.New("division by zero")
			}

			// add to operand stack
			operandsStack = append(operandsStack, calculateInfix(symbol, operand1, operand2))
		}
	}

	// Last value in stack
	return operandsStack[len(operandsStack)-1], nil
}
