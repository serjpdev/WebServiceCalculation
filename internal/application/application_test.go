package application_test

import (
	"bytes"
	"github.com/serjpdev/WebServiceCalculation/internal/application"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	testCasesSuccess := []struct {
		name             string
		path             string
		sendBody         string
		expectedBody     []byte
		expectedRespCode int
	}{
		{name: "not valid empty request",
			path:             "/api/v1/calculate",
			sendBody:         "",
			expectedBody:     []byte(`{"error": "Expression is not valid"}`),
			expectedRespCode: 422,
		},
		{name: "not valid request",
			path:             "/api/v1/calculate",
			sendBody:         `{"expression": "1+1-"}`,
			expectedBody:     []byte(`{"error": "Expression is not valid"}`),
			expectedRespCode: 422,
		},
		{name: "not valid request bracket",
			path:             "/api/v1/calculate",
			sendBody:         `{"expression": "1+1-4+((1/1)"}`,
			expectedBody:     []byte(`{"error": "Expression is not valid"}`),
			expectedRespCode: 422,
		},
		{name: "not valid request zero division",
			path:             "/api/v1/calculate",
			sendBody:         `{"expression": "1+1-4+(1/1)/0+1"}`,
			expectedBody:     []byte(`{"error": "Expression is not valid"}`),
			expectedRespCode: 422,
		},
		{name: "not valid request wrong bracket)",
			path:             "/api/v1/calculate",
			sendBody:         `{"expression": "1+1-4+(1/1("}`,
			expectedBody:     []byte(`{"error": "Expression is not valid"}`),
			expectedRespCode: 422,
		},
		{name: "valid simple1",
			path:             "/api/v1/calculate",
			sendBody:         `{"expression": "1+1"}`,
			expectedBody:     []byte(`{"result": "2.000000"}`),
			expectedRespCode: 200,
		},
		{name: "valid simple2",
			path:             "/api/v1/calculate",
			sendBody:         `{"expression": "1 / 2"}`,
			expectedBody:     []byte(`{"result": "0.500000"}`),
			expectedRespCode: 200,
		},
		{name: "valid complex",
			path:             "/api/v1/calculate",
			sendBody:         `{"expression": "2+2*2"}`,
			expectedBody:     []byte(`{"result": "6.000000"}`),
			expectedRespCode: 200,
		},
		//{name: "valid complex2",
		//	path:             "/api/v1/calculate",
		//	sendBody:         `{"expression": "(2+2)*2"}`,
		//	expectedBody:     []byte(`{"result": "8.000000"}`),
		//	expectedRespCode: 200,
		//},
		//{name: "valid complex3",
		//	path:             "/api/v1/calculate",
		//	sendBody:         `{"expression": "(42+8)*243+123+0.678"}`,
		//	expectedBody:     []byte(`{"result": "12273.678000"}`),
		//	expectedRespCode: 200,
		//},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(testCase.sendBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			application.CalcHandler(w, req)
			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != testCase.expectedRespCode {
				t.Errorf("%s: Expected %d, but got %d", testCase.name, testCase.expectedRespCode, res.StatusCode)
			}
			clearOutput := bytes.Trim(w.Body.Bytes(), " \n\t")
			if !bytes.Equal(clearOutput, testCase.expectedBody) {
				t.Errorf("%s: Expected %s but got %s", testCase.name, testCase.expectedBody, clearOutput)
			}

		})
	}
}
