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
	if component == "y" || component == "green" {
		index = 1
	} else if component == "z" || component == "blue" {
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

func isEqualAddTwoColors(tupleA, tupleB string, x, y, z float64) error {
	newTuple := fmt.Sprintf("%s + %s", tupleA, tupleB)
	tuples[newTuple] = tuples[tupleA].Add(tuples[tupleB])

	return isEqualTuple(newTuple, x, y, z, 0.0)
}

func isEqualMulTwoTuples(tupleA string, mulAmount, x, y, z, w float64) error {
	newTuple := fmt.Sprintf("%s * %f", tupleA, mulAmount)
	tuples[newTuple] = tuples[tupleA].Mul(mulAmount)

	return isEqualTuple(newTuple, x, y, z, w)
}

func isEqualColorScalarMul(tupleA string, mulAmount, x, y, z float64) error {
	newTuple := fmt.Sprintf("%s * %f", tupleA, mulAmount)
	tuples[newTuple] = tuples[tupleA].Mul(mulAmount)

	return isEqualTuple(newTuple, x, y, z, 0.0)
}

func isEqualDivTwoTuples(tupleA string, divAmount, x, y, z, w float64) error {
	newTuple := fmt.Sprintf("%s / %f", tupleA, divAmount)
	tuples[newTuple] = tuples[tupleA].Div(divAmount)

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

func isEqualDot(tupleA, tupleB string, dot float64) error {
	if gotten := tuples[tupleA].Dot(tuples[tupleB]); !compareFloat(gotten, dot) {
		return fmt.Errorf("expected %s.Dot(%s) magnitude to be %f got %f",
			tupleA, tupleB, dot, gotten)
	}
	return nil
}

func isEqualCross(tupleA, tupleB string, x, y, z float64) error {
	newTuple := fmt.Sprintf("cross(%s, %s)", tupleA, tupleB)
	tuples[newTuple] = tuples[tupleA].Cross(tuples[tupleB])

	return isEqualTuple(newTuple, x, y, z, 0.0)
}

func isEqualColorMul(tupleA, tupleB string, x, y, z float64) error {
	newTuple := fmt.Sprintf("cross(%s, %s)", tupleA, tupleB)
	tuples[newTuple] = tuples[tupleA].ColorMul(tuples[tupleB])

	return isEqualTuple(newTuple, x, y, z, 0.0)
}

func isEqualMag(tuple string, mag float64) error {
	if gotten := tuples[tuple].Magnitude(); !compareFloat(gotten, mag) {
		return fmt.Errorf("expected %s magnitude to be %f got %f", tuple, mag, gotten)
	}
	return nil
}

func isEqualNorm(tuple string, x, y, z float64) error {
	newTuple := fmt.Sprintf("normalize(%s)", tuple)
	tuples[newTuple] = tuples[tuple].Normalize()

	return isEqualTuple(newTuple, x, y, z, 0.0)
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

func createColor(tuple string, x, y, z float64) {
	tuples[tuple] = Color(x, y, z)
}

func createVectorNormalize(newTuple string, tuple string) {
	tuples[newTuple] = tuples[tuple].Normalize()
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
	ctx.Step(fmt.Sprintf(`^%s ← color\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), createVector)
	ctx.Step(fmt.Sprintf(`^%s ← normalize\(%s\)$`, wordRegex, wordRegex), createVectorNormalize)

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
	ctx.Step(fmt.Sprintf(`^%s \+ %s = color\(%s, %s, %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex),
		isEqualAddTwoColors)
	ctx.Step(fmt.Sprintf(`^%s - %s = (?:vector|color)\(%s, %s, %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex),
		isEqualSubVecRes)
	ctx.Step(fmt.Sprintf(`^%s - %s = point\(%s, %s, %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex),
		isEqualSubPointRes)
	ctx.Step(fmt.Sprintf(`^-%s = tuple\(%s, %s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex), isEqualNegate)
	ctx.Step(fmt.Sprintf(`^%s \* %s = tuple\(%s, %s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		isEqualMulTwoTuples)
	ctx.Step(fmt.Sprintf(`^%s \* %s = color\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		isEqualColorScalarMul)
	ctx.Step(fmt.Sprintf(`^%s / %s = tuple\(%s, %s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		isEqualDivTwoTuples)

	ctx.Step(fmt.Sprintf(`^magnitude\(%s\) = %s$`, wordRegex, floatRegex), isEqualMag)
	ctx.Step(fmt.Sprintf(`^normalize\(%s\) = vector\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), isEqualNorm)

	ctx.Step(fmt.Sprintf(`^dot\(%s, %s\) = %s$`,
		wordRegex, wordRegex, floatRegex),
		isEqualDot)
	ctx.Step(fmt.Sprintf(`^cross\(%s, %s\) = vector\(%s, %s, %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex),
		isEqualCross)
	ctx.Step(fmt.Sprintf(`^%s \* %s = color\(%s, %s, %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex),
		isEqualColorMul)

}
