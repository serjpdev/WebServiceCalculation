package calculation_test

import (
	"github.com/serjpdev/WebServiceCalculation/pkg/calculation"
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
		//{
		//	name:           "complex",
		//	expression:     "22+2",
		//	expectedResult: 24,
		//},
		//{
		//	name:           "complex2",
		//	expression:     "(42+8)*243+123+0.678",
		//	expectedResult: 12273.678,
		//},
		//{
		//	name:           "complex3",
		//	expression:     "(4+8) * 243+123+ 0.678",
		//	expectedResult: 12273.678,
		//},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("%s: successful case %s returns error, %s", testCase.name, testCase.expression, err)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%s: %f should be equal %f", testCase.name, val, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "simple",
			expression: "1+1*",
		},
		{
			name:       "priority",
			expression: "2+2**2",
		},
		{
			name:       "priority",
			expression: "((2+2-*(2",
		},
		{
			name:       "/",
			expression: "",
		},
		{
			name:       "div by zero",
			expression: "(2+2)/0",
		},
		{
			name:       "one more bracket",
			expression: "(2+2)/2+)(1-2)",
		},
		{
			name:       "one more bracket2",
			expression: "(2+2)/2+((2-2)",
		},
		{
			name:       "one more bracket3",
			expression: "(2+2))/2+(1-2)",
		},
		{
			name:       "one more bracket3",
			expression: "(2+2)--3",
		},
		{
			name:       "one more -",
			expression: "1+1-",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid but result  %f was obtained", testCase.expression, val)
			}
		})
	}
}
