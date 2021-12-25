package main_test

import (
	"context"
	"fmt"
	//"strconv"

	"github.com/nicholasblaskey/raytracer/matrix"
	//"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

func transformBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func transformSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`%s \* %s = point\(%s, %s, %s\)`, wordRegex, wordRegex,
		floatRegex, floatRegex, floatRegex), matrixMulPointEquals)

	ctx.Step(fmt.Sprintf(`^%s ← translation\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), createTranslation)
}

func createTranslation(t string, x, y, z float64) {
	matrices[t] = matrix.Translate(x, y, z)
}

func matrixMulPointEquals(m, p string, x, y, z float64) error {
	actual := fmt.Sprintf("%s * %s", m, p)
	tuples[actual] = matrices[m].(matrix.Mat4).Mul4x1(tuples[p])

	return isEqualTuple(actual, x, y, z, 1.0)
}
