package impl

import "testing"

func Test_Person_IsCorrectName(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"Joshua", true},
		{"joshua", false},
		{"JoshuA", false},
		{"Jos123", false},
		{"123456", false},
	}

	for _, testCase := range testCases {
		got := IsCorrectPersonName(testCase.input)
		if testCase.want != got {
			t.Errorf("Something wrong. Input %v, want %v, got %v",
				testCase.input, testCase.want, got)
		}
	}
}
