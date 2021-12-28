package main_test

import (
	"context"
	"fmt"

	//"github.com/nicholasblaskey/raytracer/ray"
	"github.com/cucumber/godog"
	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/shape"
)

var spheres map[string]*shape.Sphere
var intersections map[string][]*intersection.Intersection

func sphereBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	spheres = make(map[string]*shape.Sphere)
	intersections = make(map[string][]*intersection.Intersection)
	return ctx, nil
}

func sphereSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← sphere\(\)$`, wordRegex), createSphere)
	ctx.Step(fmt.Sprintf(`^%s ← intersect\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), intersectSphere)

	ctx.Step(fmt.Sprintf(`^%s.count = %s`, wordRegex, intRegex),
		intersectCountEqual)
	ctx.Step(fmt.Sprintf(`^%s\[%s\](?:.t|) = %s`, wordRegex, intRegex, floatRegex),
		intersectValueEqual)
	ctx.Step(fmt.Sprintf(`^%s\[%s\].object = %s`, wordRegex, intRegex, wordRegex),
		intersectObjectEqual)

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

	if intersections[i][index].T != expected {
		return fmt.Errorf("%s[%d] expected %f got %f", i, index,
			expected, intersections[i][index].T)
	}
	return nil
}

func intersectObjectEqual(i string, index int, expected string) error {
	if index >= len(intersections[i]) {
		return fmt.Errorf("Tried to get out of bounds index %d of intersection %v",
			index, len(intersections[i]))
	}

	if intersections[i][index].Obj != spheres[expected] {
		return fmt.Errorf("%s[%d] expected %v got %v", i, index,
			spheres[expected], intersections[i][index].Obj)
	}
	return nil
}
