package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"
	//"github.com/nicholasblaskey/raytracer/matrix"

	"github.com/cucumber/godog"
)

func smoothTriangleBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func smoothTriangleSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`%s ← smooth_triangle\(%s, %s, %s, %s, %s, %s\)`,
		wordRegex, wordRegex, wordRegex, wordRegex, wordRegex,
		wordRegex, wordRegex), createSmoothTriangle)

	ctx.Step(fmt.Sprintf(`%s ← normal_at\(%s, point\(%s, %s, %s\), %s\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex, wordRegex),
		normalAtWithIntersection)
}

func createSmoothTriangle(t, p0, p1, p2, n0, n1, n2 string) {
	shapes[t] = shape.NewSmoothTriangle(
		tuples[p0], tuples[p1], tuples[p2],
		tuples[n0], tuples[n1], tuples[n2])
}

func normalAtWithIntersection(n, s string, x, y, z float64, i string) {
	tuples[n] = shapes[s].NormalAt(tuple.Point(x, y, z), intersectionObjects[i])
}
