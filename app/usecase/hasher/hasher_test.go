package hasher

import (
	"testing"
	"unhashService/entity"
	"unhashService/pkg/logger"
)

type testCase struct {
	name        string
	inputHash   []entity.Hash
	inputNumber []entity.PhoneNumber
	inputDomain string
	expected    []string
	expectError bool
}

func TestHashPhoneNumber(t *testing.T) {
	logger := logger.NewConsoleLogger(logger.InfoLevel)

	tests := []testCase{
		{
			name: "valid hash",
			inputNumber: []entity.PhoneNumber{
				{PhoneNumber: "123456789012", Salt: 42},
			},
			inputDomain: "1",
			expected:    []string{"1a19181f1e1d1c13121b1a19"},
			expectError: false,
		},
		{
			name: "invalid domain",
			inputHash: []entity.Hash{
				{PhoneNumber: "123456789012", Salt: 42},
			},
			inputDomain: "invalid",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "empty input",
			inputHash:   []entity.Hash{},
			inputDomain: "1",
			expected:    []string{},
			expectError: false,
		},
	}

	uc := New(logger)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := uc.HashPhoneNumber(tc.inputNumber, tc.inputDomain)

			if (err != nil) != tc.expectError {
				t.Errorf("unexpected error status: got %v, want %v", err != nil, tc.expectError)
			}

			if !stringSlicesEqual(result, tc.expected) {
				t.Errorf("unexpected result: got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestUnhashPhoneNumber(t *testing.T) {
	logger := logger.NewConsoleLogger(logger.InfoLevel)

	tests := []testCase{
		{
			name: "valid unhash",
			inputHash: []entity.Hash{
				{PhoneNumber: "1a19181f1e1d1c13121b1a19", Salt: 42},
			},
			inputDomain: "1",
			expected:    []string{"123456789012"},
			expectError: false,
		},
		{
			name: "invalid hex",
			inputHash: []entity.Hash{
				{PhoneNumber: "invalid_hex", Salt: 42},
			},
			inputDomain: "1",
			expected:    nil,
			expectError: true,
		},
		{
			name: "invalid domain",
			inputHash: []entity.Hash{
				{PhoneNumber: "4f4e464645", Salt: 42},
			},
			inputDomain: "invalid",
			expected:    nil,
			expectError: true,
		},
	}

	uc := New(logger)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := uc.UnhashPhoneNumber(tc.inputHash, tc.inputDomain)

			if (err != nil) != tc.expectError {
				t.Errorf("unexpected error status: got %v, want %v", err != nil, tc.expectError)
			}

			if !stringSlicesEqual(result, tc.expected) {
				t.Errorf("unexpected result: got %v, want %v", result, tc.expected)
			}
		})
	}
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
