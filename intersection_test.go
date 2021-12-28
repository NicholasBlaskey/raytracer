package main_test

import (
	"context"
	"fmt"

	//"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/intersection"

	"github.com/cucumber/godog"
)

var intersectionObjects map[string]*intersection.Intersection

func intersectionBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	intersectionObjects = make(map[string]*intersection.Intersection)
	return ctx, nil
}

func intersectionSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← intersection\(%s, %s\)$`,
		wordRegex, floatRegex, wordRegex), intersectionCreate)
	// TODO figure out varadic cucumber steps.
	ctx.Step(fmt.Sprintf(`^%s ← intersections\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), aggregateIntersections)
	ctx.Step(fmt.Sprintf(`^%s ← intersections\(%s, %s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex, wordRegex),
		aggregateIntersections4)

	ctx.Step(fmt.Sprintf(`^%s ← hit\(%s\)$`, wordRegex, wordRegex),
		intersectionCreateHit)

	ctx.Step(fmt.Sprintf(`^%s\.t = %s$`, wordRegex, floatRegex),
		intersectionTimeEqual)
	ctx.Step(fmt.Sprintf(`^%s\.object = %s$`, wordRegex, wordRegex),
		intersectionObjectEqual)

	ctx.Step(fmt.Sprintf(`^%s is nothing$`, wordRegex),
		intersectionIsNothing)
}

func intersectionCreate(i string, t float64, obj string) {
	intersectionObjects[i] = intersection.New(t, spheres[obj])
}

func intersectionTimeEqual(i string, t float64) error {
	if actual := intersectionObjects[i].T; actual != t {
		return fmt.Errorf("%s.t expected %f got %f", i, t, actual)
	}
	return nil
}

func intersectionObjectEqual(i, obj string) error {
	if actual := intersectionObjects[i].Obj; actual != spheres[obj] {
		return fmt.Errorf("%s.object expected %v got %v", i, obj, actual)
	}
	return nil
}

func aggregateIntersections(res, i0, i1 string) {
	intersections[res] = intersection.Aggregate(intersectionObjects[i0],
		intersectionObjects[i1])
}

func aggregateIntersections4(res, i0, i1, i2, i3 string) {
	intersections[res] = intersection.Aggregate(intersectionObjects[i0],
		intersectionObjects[i1], intersectionObjects[i2], intersectionObjects[i3])
}

func intersectionCreateHit(hit, i string) {
	intersectionObjects[hit] = intersection.Hit(intersections[i])
}

func intersectionIsNothing(i string) error {
	if intersectionObjects[i] != nil {
		return fmt.Errorf("Expected %s to be null got %v", i, intersectionObjects[i])
	}
	return nil
}

func intersectionEquals(i0, i1 string) error {
	if intersectionObjects[i0] != intersectionObjects[i1] {
		return fmt.Errorf("expected %s (%v) to equal %s (%v)",
			i0, intersectionObjects[i0], i1, intersectionObjects[i1])
	}
	return nil
}
