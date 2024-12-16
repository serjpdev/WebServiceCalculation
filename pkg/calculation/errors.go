package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("Invalid expression")
	ErrDivisionByZero    = errors.New("division by zero")
)
