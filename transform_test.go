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
	ctx.Step(fmt.Sprintf(`%s \* %s = (point|vector)\(%s, %s, %s\)`, wordRegex, wordRegex,
		floatRegex, floatRegex, floatRegex), matrixMulVecOrPointEquals)

	ctx.Step(fmt.Sprintf(`^%s ← translation\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), createTranslation)
	ctx.Step(fmt.Sprintf(`^%s ← scaling\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), createScale)
	ctx.Step(fmt.Sprintf(`^%s ← rotation_(x|y|z)\(%s\)$`,
		wordRegex, floatRegex), createRotatation)
	ctx.Step(fmt.Sprintf(`^%s ← shearing\(%s, %s, %s, %s, %s, %s\)$`, wordRegex,
		floatRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		createShear)
}

func createTranslation(t string, x, y, z float64) {
	matrices[t] = matrix.Translate(x, y, z)
}

func createScale(t string, x, y, z float64) {
	matrices[t] = matrix.Scale(x, y, z)
}

func matrixMulVecOrPointEquals(m, p, pointOrVector string, x, y, z float64) error {
	actual := fmt.Sprintf("%s * %s", m, p)
	tuples[actual] = matrices[m].(matrix.Mat4).Mul4x1(tuples[p])

	if pointOrVector == "point" {
		return isEqualTuple(actual, x, y, z, 1.0)
	}
	return isEqualTuple(actual, x, y, z, 0.0)
}

func createRotatation(t, xyz string, theta float64) {
	if xyz == "x" {
		matrices[t] = matrix.RotateX(theta)
	} else if xyz == "y" {
		matrices[t] = matrix.RotateY(theta)
	} else {
		matrices[t] = matrix.RotateZ(theta)
	}
}

func createShear(t string, xy, xz, yx, yz, zx, zy float64) {
	matrices[t] = matrix.Shear(xy, xz, yx, yz, zx, zy)
}
