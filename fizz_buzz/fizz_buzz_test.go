package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input  int
		output string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{6, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, tt := range tests {
		s := FizzBuzz(tt.input)
		if s != tt.output {
			t.Errorf("FizzBuzz(%v)が%vじゃなくて%qだよ", tt.input, tt.output, s)
		}
	}
}
