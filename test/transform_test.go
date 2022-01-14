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
	ctx.Step(fmt.Sprintf(`%s \* %s = (point|vector)\(%s, %s, %s\)$`, wordRegex, wordRegex,
		floatRegex, floatRegex, floatRegex), matrixMulVecOrPointEquals)
	ctx.Step(fmt.Sprintf(`%s ← %s \* %s \* %s$`, wordRegex, wordRegex,
		wordRegex, wordRegex), matrixMulMatrixMulMatrix)

	ctx.Step(fmt.Sprintf(`^%s ← translation\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), createTranslation)
	ctx.Step(fmt.Sprintf(`^%s ← scaling\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), createScale)
	ctx.Step(fmt.Sprintf(`^%s ← rotation_(x|y|z)\(%s\)$`,
		wordRegex, floatRegex), createRotatation)
	ctx.Step(fmt.Sprintf(`^%s ← shearing\(%s, %s, %s, %s, %s, %s\)$`, wordRegex,
		floatRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		createShear)
	ctx.Step(fmt.Sprintf(`^%s ← view_transform\(%s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex), createView)

	ctx.Step(fmt.Sprintf(`^%s ← scaling\(%s, %s, %s\) \* rotation_z\(%s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex), createScaleMulRotZ)
	ctx.Step(fmt.Sprintf(`^%s = scaling\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), matrixEqualToScaling)
	ctx.Step(fmt.Sprintf(`^%s = translation\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), matrixEqualToTranslation)

	ctx.Step(fmt.Sprintf(`^set_transform\(%s, rotation_y\(%s\)\)$`,
		wordRegex, floatRegex), setTransformToRotY)
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

func matrixMulMatrixMulMatrix(res string, t0, t1, t2 string) {
	matrices[res] = matrices[t0].(matrix.Mat4).Mul4(
		matrices[t1].(matrix.Mat4)).Mul4(
		matrices[t2].(matrix.Mat4))
}

func createScaleMulRotZ(t string, x, y, z, theta float64) {
	matrices[t] = matrix.Scale(x, y, z).Mul4(matrix.RotateZ(theta))
}

func createView(t, from, to, up string) {
	matrices[t] = matrix.View(tuples[from], tuples[to], tuples[up])
}

func matrixEqualToScaling(m string, x, y, z float64) error {
	expected := fmt.Sprintf("scaling(%f, %f, %f)", x, y, z)
	createScale(expected, x, y, z)
	return matrixEquals(m, expected)
}

func matrixEqualToTranslation(m string, x, y, z float64) error {
	expected := fmt.Sprintf("translation(%f, %f, %f)", x, y, z)
	createTranslation(expected, x, y, z)
	return matrixEquals(m, expected)
}

func setTransformToRotY(s string, theta float64) {
	shapes[s].SetTransform(matrix.RotateY(theta))
}
