package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/shape"

	"github.com/cucumber/godog"
)

func triangleBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func triangleSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← triangle\(%s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex), createTriangle)

	ctx.Step(fmt.Sprintf(`^%s.(p1|p2|p3) = %s$`, wordRegex, wordRegex),
		trianglePointEqualTo)
	ctx.Step(fmt.Sprintf(`^%s.(e1|e2|normal) = vector\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), triangleVectorEqualTo)
}

func createTriangle(t, p0, p1, p2 string) {
	shapes[t] = shape.NewTriangle(tuples[p0], tuples[p1], tuples[p2])
}

func trianglePointEqualTo(t, whichPoint, expected string) error {
	tri := shapes[t].(*shape.Triangle)
	actual := tri.P0
	if whichPoint == "p2" {
		actual = tri.P1
	} else if whichPoint == "p3" {
		actual = tri.P2
	}

	return isEqualTuple(expected, actual[0], actual[1], actual[2], actual[3])
}

func triangleVectorEqualTo(t, whichVec string, x, y, z float64) error {
	tri := shapes[t].(*shape.Triangle)
	actual := fmt.Sprintf("actual %s.%s", t, whichVec)
	if whichVec == "e1" {
		tuples[actual] = tri.E0
	} else if whichVec == "e2" {
		tuples[actual] = tri.E1
	} else {
		tuples[actual] = tri.Normal
	}

	return isEqualTuple(actual, x, y, z, tuples[actual][3])
}
