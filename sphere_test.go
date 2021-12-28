package main_test

import (
	"context"
	"fmt"

	//"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/shape"

	"github.com/cucumber/godog"
)

var spheres map[string]*shape.Sphere
var intersections map[string][]float64

func sphereBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	spheres = make(map[string]*shape.Sphere)
	intersections = make(map[string][]float64)
	return ctx, nil
}

func sphereSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← sphere\(\)$`, wordRegex), createSphere)
	ctx.Step(fmt.Sprintf(`^%s ← intersect\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), intersectSphere)

	ctx.Step(fmt.Sprintf(`^%s.count = %s`, wordRegex, intRegex),
		intersectCountEqual)
	ctx.Step(fmt.Sprintf(`^%s\[%s\] = %s`, wordRegex, intRegex, floatRegex),
		intersectValueEqual)
}

func createSphere(s string) {
	spheres[s] = shape.NewSphere()
}

func intersectSphere(intersection, s, r string) {
	intersections[intersection] = spheres[s].Intersections(rays[r])
}

func intersectCountEqual(i string, expected int) error {
	actual := len(intersections[i])
	if actual != expected {
		return fmt.Errorf("%s.count expected %d got %d", i, expected, actual)
	}
	return nil
}

func intersectValueEqual(i string, index int, expected float64) error {
	if index >= len(intersections[i]) {
		return fmt.Errorf("Tried to get out of bounds index %d of intersection %v",
			index, len(intersections[i]))
	}

	if intersections[i][index] != expected {
		return fmt.Errorf("%s[%d] expected %f got %f", i, index,
			expected, intersections[i])
	}
	return nil
}
