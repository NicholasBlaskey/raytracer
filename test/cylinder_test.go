package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/shape"

	"github.com/cucumber/godog"
)

func cylinderBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func cylinderSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`%s ← cylinder\(\)`, wordRegex), createCylinder)
}

func createCylinder(s string) {
	shapes[s] = shape.NewCylinder()
}
