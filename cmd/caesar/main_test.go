package main

import "testing"

type testCase struct {
	input    string
	shift    int
	expected string
}

func TestEncrypt(t *testing.T) {
	testCases := []*testCase{
		{"abcdefghijklmnopqrstuvwxyz1234567890", 2, "cdefghijklmnopqrstuvwxyzab1234567890"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 2, "CDEFGHIJKLMNOPQRSTUVWXYZAB1234567890"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 26, "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"},
	}

	for _, testCase := range testCases {
		encrypted := encrypt(testCase.input, testCase.shift)
		if encrypted != testCase.expected {
			t.Fail()
		}
	}
}
