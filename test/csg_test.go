package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/shape"

	"github.com/cucumber/godog"
)

var booleans map[string]bool

func csgBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	booleans = make(map[string]bool)
	return ctx, nil
}

func csgSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← csg\("(union|intersection|difference)", %s, %s\)$`,
		wordRegex, wordRegex, wordRegex), createCSG)
	ctx.Step(fmt.Sprintf(`^%s ← csg\("(union|intersection|difference)", sphere\(\), cube\(\)\)$`,
		wordRegex), createCSGSphereCube)

	ctx.Step(fmt.Sprintf(`^%s\.operation = "(union|intersection|difference)"$`,
		wordRegex), csgOperationEqualTo)
	ctx.Step(fmt.Sprintf(`^%s\.(left|right) = %s$`, wordRegex, wordRegex),
		csgChildEqualTo)

	ctx.Step(fmt.Sprintf(`^%s ← intersection_allowed\("%s", %s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex, wordRegex), intersectionAllowed)
	ctx.Step(fmt.Sprintf(`^%s ← filter_intersections\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), filterIntersections)

	ctx.Step(fmt.Sprintf(`^%s\[%s\] = %s\[%s\]$`,
		wordRegex, intRegex, wordRegex, intRegex), intersectionCompareArrayElement)
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

func createCSGSphereCube(res, operation string) {
	opt := shape.Union
	if operation == "intersection" {
		opt = shape.Intersection
	} else if operation == "difference" {
		opt = shape.Difference
	}

	shapes[res] = shape.NewCSG(opt, shape.NewSphere(), shape.NewCube())
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

func intersectionAllowed(res, operation, lHitS, inLS, inRS string) {
	opt := shape.Union
	if operation == "intersection" {
		opt = shape.Intersection
	} else if operation == "difference" {
		opt = shape.Difference
	}

	lHit := lHitS == "true"
	inL := inLS == "true"
	inR := inRS == "true"

	booleans[res] = shape.IntersectionAllowed(opt, lHit, inL, inR)
}

func booleansEqual(actual string, expected bool) error {
	if b, ok := booleans[actual]; ok {
		if b != expected {
			return fmt.Errorf("%s expected %t got %t", actual, expected, b)
		}
		return nil
	} else {
		return fmt.Errorf("%s not set", actual)
	}
}

func filterIntersections(res, csg, xs string) {
	intersections[res] = shapes[csg].(*shape.CSG).FilterIntersections(intersections[xs])
}

func intersectionCompareArrayElement(xs0 string, i0 int, xs1 string, i1 int) error {
	actual := intersections[xs0][i0]
	expected := intersections[xs1][i1]

	if actual != expected {
		return fmt.Errorf("%s[%d] expected %+v got %+v", xs0, i0, expected, actual)
	}

	return nil
}
