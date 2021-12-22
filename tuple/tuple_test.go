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
	wordRegex  = `([A-Za-z0-9^\s]+)`
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

func isEqualTuple(tuple string, x, y, z, w float64) error {
	if !(compareFloat(x, tuples[tuple][0]) &&
		compareFloat(y, tuples[tuple][1]) &&
		compareFloat(z, tuples[tuple][2]) &&
		compareFloat(w, tuples[tuple][3])) {
		return fmt.Errorf("expected %s (%f, %f, %f, %f) to equal (%f, %f, %f, %f)",
			tuple, x, y, z, w,
			tuples[tuple][0], tuples[tuple][1], tuples[tuple][2], tuples[tuple][3])
	}
	return nil
}

func isEqualNegate(tuple string, x, y, z, w float64) error {
	newTuple := "-" + tuple
	tuples[newTuple] = tuples[tuple].Negate()

	return isEqualTuple(newTuple, x, y, z, w)
}

func isEqualAddTwoTuples(tupleA, tupleB string, x, y, z, w float64) error {
	newTuple := fmt.Sprintf("%s + %s", tupleA, tupleB)
	tuples[newTuple] = tuples[tupleA].Add(tuples[tupleB])

	return isEqualTuple(newTuple, x, y, z, w)
}

func isEqualSubVecRes(tupleA, tupleB string, x, y, z float64) error {
	newTuple := fmt.Sprintf("%s - %s", tupleA, tupleB)
	tuples[newTuple] = tuples[tupleA].Sub(tuples[tupleB])

	return isEqualTuple(newTuple, x, y, z, 0.0)
}

func isEqualSubPointRes(tupleA, tupleB string, x, y, z float64) error {
	newTuple := fmt.Sprintf("%s - %s", tupleA, tupleB)
	tuples[newTuple] = tuples[tupleA].Sub(tuples[tupleB])

	return isEqualTuple(newTuple, x, y, z, 1.0)
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

func createVector(tuple string, x, y, z float64) {
	tuples[tuple] = Vector(x, y, z)
}

func createPoint(tuple string, x, y, z float64) {
	tuples[tuple] = Point(x, y, z)
}

func initializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tuples = make(map[string]Tuple)

		return ctx, nil
	})

	// Create tuples.
	ctx.Step(fmt.Sprintf(`^%s ← tuple\(%s, %s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex), createTuple)
	ctx.Step(fmt.Sprintf(`^%s ← point\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), createPoint)
	ctx.Step(fmt.Sprintf(`^%s ← vector\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), createVector)

	// Check tuples are.
	ctx.Step(fmt.Sprintf(`^%s\.%s = %s$`, wordRegex, wordRegex, floatRegex),
		tupleValueEqual)
	ctx.Step(fmt.Sprintf(`^%s (is|is not) a (point|vector)$`, wordRegex), tupleIsA)
	ctx.Step(fmt.Sprintf(`^%s = tuple\(%s, %s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex), isEqualTuple)

	// Operations.
	ctx.Step(fmt.Sprintf(`^%s \+ %s = tuple\(%s, %s, %s, %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		isEqualAddTwoTuples)
	ctx.Step(fmt.Sprintf(`^%s - %s = vector\(%s, %s, %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex),
		isEqualSubVecRes)
	ctx.Step(fmt.Sprintf(`^%s - %s = point\(%s, %s, %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex),
		isEqualSubPointRes)
	ctx.Step(fmt.Sprintf(`^-%s = tuple\(%s, %s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex), isEqualNegate)
}
