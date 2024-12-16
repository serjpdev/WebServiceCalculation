package application_test

import (
	"bytes"
	"github.com/poserj/calc_project/internal/application"
	"io"
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
		{name: "valid simple",
			path:             "/api/v1/calculate",
			sendBody:         `{"expression": "1+1"}`,
			expectedBody:     []byte(`{"result": "2.000000"}`),
			expectedRespCode: 200,
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(testCase.sendBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			application.CalcHandler(w, req)
			res := w.Result()
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			strData := strings.TrimSpace(string(data))
			if res.StatusCode != testCase.expectedRespCode {
				t.Errorf("Expected %d, but got %d", testCase.expectedRespCode, res.StatusCode)
			}
			if bytes.Equal(w.Body.Bytes(), testCase.expectedBody) {
				t.Errorf("Expected %s but got %v", testCase.expectedBody, strData)
			}

		})
	}
}
