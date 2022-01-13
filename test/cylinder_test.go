package main_test

import (
	"context"
	"fmt"
	"math"

	"github.com/nicholasblaskey/raytracer/shape"

	"github.com/cucumber/godog"
)

func cylinderBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func cylinderSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← cylinder\(\)$`, wordRegex), createCylinder)
	ctx.Step(fmt.Sprintf(`^%s ← cone\(\)$`, wordRegex), createCone)

	ctx.Step(fmt.Sprintf(`^%s.(minimum|maximum) ← %s$`,
		wordRegex, floatRegex), cylinderAssignBounds)
	ctx.Step(fmt.Sprintf(`^%s.(minimum|maximum) = (-infinity|infinity)$`,
		wordRegex), cylinderBoundsEqualTo)

	ctx.Step(fmt.Sprintf(`^%s.closed ← (true|false)$`,
		wordRegex), cylinderAssignClosed)
	ctx.Step(fmt.Sprintf(`^%s.closed = (true|false)$`,
		wordRegex), cylinderClosedEqualTo)
}

func createCone(s string) {
	shapes[s] = shape.NewCone()
}

func createCylinder(s string) {
	shapes[s] = shape.NewCylinder()
}

func cylinderAssignBounds(s, minOrMax string, v float64) {
	if minOrMax == "minimum" {
		shapes[s].(*shape.Cylinder).Min = v
	} else {
		shapes[s].(*shape.Cylinder).Max = v
	}
}

func cylinderBoundsEqualTo(s, minOrMax, posOrNegInf string) error {
	expected := math.Inf(1)
	if posOrNegInf == "-infinity" {
		expected = math.Inf(-1)
	}

	actual := shapes[s].(*shape.Cylinder).Max
	if minOrMax == "minimum" {
		actual = shapes[s].(*shape.Cylinder).Min
	}

	if expected != actual {
		return fmt.Errorf("%s.%s expected %f got %f", s, minOrMax, expected, actual)
	}
	return nil
}

func cylinderAssignClosed(s, isClosed string) {
	shapes[s].(*shape.Cylinder).Closed = isClosed == "true"
}

func cylinderClosedEqualTo(s, isClosed string) error {
	expected := isClosed == "true"
	actual := shapes[s].(*shape.Cylinder).Closed

	if expected != actual {
		return fmt.Errorf("%s.Closed expected %t, got %t", s, expected, actual)
	}
	return nil
}
