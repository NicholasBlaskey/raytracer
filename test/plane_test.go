package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

func planeBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func planeSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← plane\(\)$`, wordRegex), createPlane)
	ctx.Step(fmt.Sprintf(`^%s ← local_normal_at\(%s, point\(%s, %s, %s\)\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex), localNormalAt)
}

func createPlane(p string) {
	shapes[p] = shape.NewPlane()
}

// TODO make this actually test local normal at.
func localNormalAt(n, p string, x, y, z float64) {
	if shapes[p].GetTransform() != matrix.Ident4() {
		panic("Does not support non ident transforms yet")
	}

	tuples[n] = shapes[p].NormalAt(tuple.Point(x, y, z))
}
