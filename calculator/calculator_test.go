package calculator

import "testing"

func TestSum(t *testing.T) {

	t.Log("testing calculator.Sum...")

	sum := Sum(1, 1)

	if sum != 2 {
		t.Fatalf("failed calculation. The expected result should be %v , but it was %v", 2, sum)
	}

}
