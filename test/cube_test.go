package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/shape"
	//"github.com/nicholasblaskey/raytracer/matrix"

	"github.com/cucumber/godog"
)

func cubeBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func cubeSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`%s ← cube\(\)`, wordRegex), createCube)
}

func createCube(c string) {
	shapes[c] = shape.NewCube()
}
