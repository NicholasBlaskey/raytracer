package main_test

import (
	"context"
	"fmt"

	//"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/intersection"

	"github.com/cucumber/godog"
)

var intersectionObjects map[string]intersection.Intersection

func intersectionBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	intersectionObjects = make(map[string]intersection.Intersection)
	return ctx, nil
}

func intersectionSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← intersection\(%s, %s\)$`,
		wordRegex, floatRegex, wordRegex), intersectionCreate)

	ctx.Step(fmt.Sprintf(`^%s\.t = %s$`, wordRegex, floatRegex),
		intersectionTimeEqual)
	ctx.Step(fmt.Sprintf(`^%s\.object = %s$`, wordRegex, wordRegex),
		intersectionObjectEqual)
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
