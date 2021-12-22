package tuple

import (
	"context"
	"fmt"
	"math"

	"testing"

	"github.com/cucumber/godog"
)

const (
	epsilon    = 0.00001
	wordRegex  = `([^\s]+)`
	floatRegex = `(\-*\d+\.\d+)`
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: initializeScenario,

		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"."},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

/*
func aIsAPoint() error {
	return godog.ErrPending
}

func aIsNotAVector() error {
	return godog.ErrPending
}

func aTuple(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 int) error {
	return godog.ErrPending
}

func aw(arg1, arg2 int) error {
	return godog.ErrPending
}

func ax(arg1, arg2 int) error {
	return godog.ErrPending
}

func ay(arg1, arg2 int) error {
	return godog.ErrPending
}

func az(arg1, arg2 int) error {
	return godog.ErrPending
}
*/

// TODO make this a context type!
var tuples map[string]Tuple

func createTuple(tuple string, x, y, z, w float64) {
	tuples[tuple] = Tuple{x, y, z, w}
}

func compareFloat(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func tupleValueEqual(tuple string, component string, v float64) error {
	index := 0
	if component == "y" {
		index = 1
	} else if component == "z" {
		index = 2
	} else if component == "w" {
		index = 3
	}

	if !compareFloat(v, tuples[tuple][index]) {
		return fmt.Errorf("expected %s to be %f got %f instead",
			tuple, v, tuples[tuple][index])
	}
	return nil
}

func tupleIsA(tuple string, isOrNotIs string, expected string) error {
	actual := "vector"
	if tuples[tuple][3] == 1 {
		actual = "point"
	}

	if actual == expected {
		if isOrNotIs == "not is" {
			return fmt.Errorf("expected %s (a %s) to not be a %s", tuple, actual, expected)
		}
	} else {
		if isOrNotIs == "is" {
			return fmt.Errorf("expected %s (a %s) to be a %s", tuple, actual, expected)
		}
	}
	return nil
}

func initializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tuples = make(map[string]Tuple)

		return ctx, nil
	})

	ctx.Step(fmt.Sprintf(`^%s ← tuple\(%s, %s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex), createTuple)
	ctx.Step(fmt.Sprintf(`^%s\.%s = %s$`, wordRegex, wordRegex, floatRegex),
		tupleValueEqual)
	ctx.Step(fmt.Sprintf(`^%s (is|is not) a (point|vector)$`, wordRegex), tupleIsA)

	//ctx.Step("^\s is a point$", isAPoint)
	/*
		ctx.Step(`^a is a point$`, aIsAPoint)
		ctx.Step(`^a is not a vector$`, aIsNotAVector)

		ctx.Step(`^a\.w = (\d+)\.(\d+)$`, aw)
		ctx.Step(`^a\.x = (\d+)\.(\d+)$`, ax)
		ctx.Step(`^a\.y = -(\d+)\.(\d+)$`, ay)
		ctx.Step(`^a\.z = (\d+)\.(\d+)$`, az)
	*/
}
