package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

var tuples map[string]tuple.Tuple

func tuplesBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	tuples = make(map[string]tuple.Tuple)
	return ctx, nil
}

func tuplesSteps(ctx *godog.ScenarioContext) {
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
	ctx.Step(fmt.Sprintf(`^%s\.(x|y|z|w|red|green|blue) = %s$`, wordRegex, floatRegex),
		tupleValueEqual)
	ctx.Step(fmt.Sprintf(`^%s (is|is not) a (point|vector)$`, wordRegex), tupleIsA)
	ctx.Step(fmt.Sprintf(`^%s = tuple\(%s, %s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex), isEqualTuple)
	ctx.Step(fmt.Sprintf(`^%s = point\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), isEqualPoint)

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

func createTuple(t string, x, y, z, w float64) {
	tuples[t] = tuple.Tuple{x, y, z, w}
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

func isEqualPoint(tuple string, x, y, z float64) error {
	return isEqualTuple(tuple, x, y, z, 1.0)
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

func createVector(t string, x, y, z float64) {
	tuples[t] = tuple.Vector(x, y, z)
}

func createColor(t string, x, y, z float64) {
	tuples[t] = tuple.Color(x, y, z)
}

func createVectorNormalize(newTuple string, t string) {
	tuples[newTuple] = tuples[t].Normalize()
}

func createPoint(t string, x, y, z float64) {
	tuples[t] = tuple.Point(x, y, z)
}
