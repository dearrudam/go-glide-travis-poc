package calculator

import (
	"fmt"
	"os"
	"testing"
	"time"
	"github.com/DATA-DOG/godog"
)


func TestMain(m *testing.M) {
	format := "pretty"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}

	status := godog.RunWithOptions("calculator", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: format,
		Paths:     []string{"features"},
		StopOnFailure: true,
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

var result int

func iSumWith(arg1, arg2 int) error {

	result=Sum(arg1,arg2)
	return nil
}

func theResultShouldBe(arg1 int) error {
	if result !=arg1 {
		return fmt.Errorf("invalid calculation: expected %d, but it was %d", arg1, result)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I sum (\d+) with (\d+)$`, iSumWith)
	s.Step(`^the result should be (\d+)$`, theResultShouldBe)
}

func TestSum(t *testing.T) {

	t.Log("testing calculator.Sum...")

	sum := Sum(1, 1)

	if sum != 2 {
		t.Fatalf("failed calculation. The expected result should be %v , but it was %v", 2, sum)
	}

}
