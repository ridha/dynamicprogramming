package dynamicprogramming

import "testing"

func TestCombinationSum4(t *testing.T) {
	actual := CombinationSum4([]int{1, 2, 3}, 4)
	expected := 7
	if expected != actual {
		t.Errorf("Expected: %+v, Actual: %+v", expected, actual)
	}

}
