package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/shape"

	"github.com/cucumber/godog"
)

func csgBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func csgSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← csg\("(union|intersection|difference)", %s, %s\)$`,
		wordRegex, wordRegex, wordRegex), createCSG)

	ctx.Step(fmt.Sprintf(`^%s\.operation = "(union|intersection|difference)"$`,
		wordRegex), csgOperationEqualTo)
	ctx.Step(fmt.Sprintf(`^%s\.(left|right) = %s$`, wordRegex, wordRegex),
		csgChildEqualTo)
}

func createCSG(res, operation, left, right string) {
	opt := shape.Union
	if operation == "intersection" {
		opt = shape.Intersection
	} else if operation == "difference" {
		opt = shape.Difference
	}

	shapes[res] = shape.NewCSG(opt, shapes[left], shapes[right])
}

func csgOperationEqualTo(c, operation string) error {
	csg := shapes[c].(*shape.CSG)
	expected := "union"
	if csg.Operation == shape.Intersection {
		expected = "intersection"
	} else if csg.Operation == shape.Difference {
		expected = "difference"
	}

	if expected != operation {
		return fmt.Errorf("%s.operation expected %s got %s", c, expected, operation)
	}
	return nil
}

func csgChildEqualTo(c, leftOrRight, expected string) error {
	csg := shapes[c].(*shape.CSG)
	actual := csg.Left
	if leftOrRight == "right" {
		actual = csg.Right
	}

	if shapes[expected] != actual {
		return fmt.Errorf("%s.%s expected %+v got %+v", c, leftOrRight,
			shapes[expected], actual)
	}
	return nil
}
