package test

import (
	"tdd/fn"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		num         string
		expected    int
		description string
	}{
		{
			description: "test_case_1",
			num:         "1,3",
			expected:    4,
		},
		{
			description: "test_case_2",
			num:         "1\n2,3",
			expected:    6,
		}, {
			description: "test_case_3",
			num:         "//;\n1;2",
			expected:    3,
		}, {
			description: "test_case_4",
			num:         "//;3;11;3",
			expected:    17,
		}, {
			description: "test_case_5",
			num:         "//;1;1;3",
			expected:    5,
		},
		{
			description: "test_case_6",
			num:         "//;",
			expected:    0,
		}, {
			description: "test_case_6",
			num:         "//;4",
			expected:    4,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result, er := fn.Add(tt.num)
			if er != nil {
				t.Errorf("expected %s", er)
			}
			if result != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
