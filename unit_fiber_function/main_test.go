package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 1, 2, 3},
		{"Add negative numbers", 1, -3, -2},
		{"Add zero", 0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Add(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("Add(%d, %d): expected %d, actual %d", tc.a, tc.b, tc.expected, actual)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name     string
		number   int
		expected int
	}{
		{"Factorial of 0", 0, 1},
		{"Factorial of 1", 1, 1},
		{"Factorial of 4", 4, 24},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Factorial(tc.number)

			if actual != tc.expected {
				t.Errorf("Factorial(%d): expected %d, actual %d", tc.number, tc.expected, actual)
			}
		})
	}
}

func TestUserRoute(t *testing.T) {
	app := setup()

	tests := []struct {
		description  string
		requestBody  User
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  User{"jane.doe@example.com", "Jane Doe", 30},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Invalid email",
			requestBody:  User{"invalid-email", "Jane Doe", 30},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid fullname",
			requestBody:  User{"jane.doe@example.com", "12345", 30},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid age",
			requestBody:  User{"jane.doe@example.com", "Jane Doe", -5},
			expectStatus: fiber.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
}
