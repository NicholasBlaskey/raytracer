package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

var rays map[string]ray.Ray

func rayBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	rays = make(map[string]ray.Ray)
	return ctx, nil
}

func raySteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← ray\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), createRay)
	ctx.Step(fmt.Sprintf(`^%s ← ray\(point\(%s, %s, %s\), %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, wordRegex),
		createRayFromHalfLiteral)
	ctx.Step(fmt.Sprintf(`^%s ← ray\(point\(%s, %s, %s\), vector\(%s, %s, %s\)\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex,
		floatRegex, floatRegex, floatRegex), createRayFromLiteral)
	ctx.Step(fmt.Sprintf(`^%s ← transform\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), transformRay)

	ctx.Step(fmt.Sprintf(`^%s.(origin|direction) = %s$`,
		wordRegex, wordRegex), rayComponentEqual)
	ctx.Step(fmt.Sprintf(`^%s.(origin|direction) = (vector|point)\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), rayComponentEqualVecOrPoint)
	ctx.Step(fmt.Sprintf(`^position\(%s, %s\) = point\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex), rayPosAtEquals)
}

func createRay(r, origin, dir string) {
	rays[r] = ray.New(tuples[origin], tuples[dir])
}

func createRayFromHalfLiteral(r string, ox, oy, oz float64, dir string) {
	rays[r] = ray.New(tuple.Point(ox, oy, oz), tuples[dir])
}

func createRayFromLiteral(r string, ox, oy, oz, dx, dy, dz float64) {
	rays[r] = ray.New(tuple.Point(ox, oy, oz), tuple.Vector(dx, dy, dz))
}

func rayComponentEqual(r, originOrDir, equalTo string) error {
	actual := rays[r].Direction
	if originOrDir == "origin" {
		actual = rays[r].Origin
	}

	return isEqualTuple(equalTo, actual[0], actual[1], actual[2], actual[3])
}

func rayComponentEqualVecOrPoint(r, originOrDir,
	vecOrPoint string, x, y, z float64) error {

	expected := fmt.Sprintf("%s.%s", r, originOrDir)
	if vecOrPoint == "vector" {
		tuples[expected] = tuple.Vector(x, y, z)
	} else {
		tuples[expected] = tuple.Point(x, y, z)
	}

	return rayComponentEqual(r, originOrDir, expected)
}

func rayPosAtEquals(r string, t, x, y, z float64) error {
	actual := fmt.Sprintf(`position(%s, %f)`, r, t)
	tuples[actual] = rays[r].PositionAt(t)
	return isEqualTuple(actual, x, y, z, 1.0)
}

func transformRay(res, r, transformation string) {
	rays[res] = rays[r].Transform(matrices[transformation].(matrix.Mat4))
}
