package dynamicprogramming

import "testing"

func TestCoinChange(t *testing.T) {
	actual := CoinChange([]int{1, 2, 5}, 11)
	expected := 3
	if expected != actual {
		t.Errorf("Expected: %+v, Actual: %+v", expected, actual)
	}

	actual = CoinChange([]int{2}, 3)
	expected = -1
	if expected != actual {
		t.Errorf("Expected: %+v, Actual: %+v", expected, actual)
	}

	actual = CoinChange([]int{1}, 1)
	expected = 1
	if expected != actual {
		t.Errorf("Expected: %+v, Actual: %+v", expected, actual)
	}

}
