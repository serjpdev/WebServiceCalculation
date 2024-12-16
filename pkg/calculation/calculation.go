package calculation

import (
	"errors"
	"strconv"
	"unicode"
)

type MathOp struct{}

func (s *MathOp) operate(b float64, operator rune, a float64) (float64, error) {
	switch operator {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0.0, errors.New("Division by zero")
		}
		return a / b, nil
	default:
		return 0, nil
	}
}

func (s *MathOp) calculate(expression string) (float64, error) {
	operatorPrecedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'(': 0,
		')': 0,
	}

	operatorStack := []rune{}
	operandStack := []float64{}
	stringLength := len(expression)
	i := 0

	// Obhod string
	for i < stringLength {
		char := rune(expression[i])
		if char == ' ' {
			i++
			continue
		}
		if unicode.IsDigit(char) || char == '.' {
			curDigiteStr := ""
			isDot := false

			for i < stringLength && (unicode.IsDigit(rune(expression[i])) || rune(expression[i]) == '.') {
				isDot = expression[i] == '.'
				curDigiteStr += string(expression[i])
				i++
			}
			if !isDot {
				curDigiteStr += ".0"
			}
			realNum, _ := strconv.ParseFloat(curDigiteStr, 64)
			isDot = false
			operandStack = append(operandStack, realNum)
			continue
		}

		if len(operatorStack) == 0 || char == '(' || operatorPrecedence[char] > operatorPrecedence[operatorStack[len(operatorStack)-1]] {
			if len(operandStack) == 0 && (char == '-' || char == '+') {
				operandStack = append(operandStack, 0)
			}
			operatorStack = append(operatorStack, char)
			if char == '(' {
				j := i
				for j+1 < stringLength {

					if expression[j+1] == '-' || expression[j+1] == '+' {
						operandStack = append(operandStack, 0)
					}
					if expression[j+1] != ' ' {
						break
					}
					j++
				}
			}
		} else if char == ')' {
			op := operatorStack[len(operatorStack)-1]
			operatorStack = operatorStack[:len(operatorStack)-1]
			for op != '(' {
				a := operandStack[len(operandStack)-1]
				operandStack = operandStack[:len(operandStack)-1]
				b := operandStack[len(operandStack)-1]
				operandStack = operandStack[:len(operandStack)-1]
				tmp, err := s.operate(a, op, b)
				if err != nil {
					return 0, err
				}
				operandStack = append(operandStack, tmp)
				op = operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
		} else if operatorPrecedence[char] <= operatorPrecedence[operatorStack[len(operatorStack)-1]] {

			op := operatorStack[len(operatorStack)-1]
			for len(operatorStack) > 0 && operatorPrecedence[char] <= operatorPrecedence[operatorStack[len(operatorStack)-1]] && op != '(' {
				operatorStack = operatorStack[:len(operatorStack)-1]
				a := operandStack[len(operandStack)-1]
				operandStack = operandStack[:len(operandStack)-1]
				b := operandStack[len(operandStack)-1]
				operandStack = operandStack[:len(operandStack)-1]
				tmp, err := s.operate(a, op, b)
				if err != nil {
					return 0, err
				}
				operandStack = append(operandStack, tmp)
				if len(operatorStack) > 0 {
					op = operatorStack[len(operatorStack)-1]
				} else {
					break
				}
			}
			operatorStack = append(operatorStack, char)
		}
		i++
	}

	for len(operatorStack) > 0 {
		lenOperandStack := len(operandStack)
		if len(operandStack) <= 1 {
			return 0, errors.New("Not valid string")
		}
		op := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		a := operandStack[lenOperandStack-1]
		operandStack = operandStack[:len(operandStack)-1]
		b := operandStack[len(operandStack)-1]
		operandStack = operandStack[:len(operandStack)-1]
		tmp, err := s.operate(a, op, b)
		if err != nil {
			return 0, err
		}
		operandStack = append(operandStack, tmp)
	}

	return operandStack[len(operandStack)-1], nil
}

func Calc(expression string) (float64, error) {
	if len(expression)%2 == 0 {
		return 0, errors.New("not valid")

	}
	mathOp := MathOp{}
	result, err := mathOp.calculate(expression)
	return result, err
}
