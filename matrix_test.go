package main_test

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nicholasblaskey/raytracer/matrix"
	//"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

var matrices map[string]matrix.Matrix

func matrixBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	matrices = make(map[string]matrix.Matrix)

	return ctx, nil
}

func matrixSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^the following %sx%s matrix %s:$`,
		intRegex, intRegex, wordRegex), matrixCreate)
	ctx.Step(fmt.Sprintf(`^the following matrix %s:$`, wordRegex), matrixCreate4)
	ctx.Step(fmt.Sprintf(`^%s\[%s,%s\] = %s$`,
		wordRegex, intRegex, intRegex, floatRegex), matrixElementEqual)
	ctx.Step(fmt.Sprintf(`^%s = %s$`,
		wordRegex, wordRegex), matrixEquals)
	ctx.Step(fmt.Sprintf(`^%s != %s$`,
		wordRegex, wordRegex), matrixNotEquals)

}

func matrixElementEqual(mat string, i, j int, expected float64) error {
	if actual := matrices[mat].At(i, j); actual != expected {
		return fmt.Errorf("For %s[%d,%d] expected %f got %f", mat, i, j, expected, actual)
	}
	return nil
}

func matrixCreate(n, m int, mat string, data *godog.Table) error {
	rows := [][]float64{}
	for _, row := range data.Rows {
		r := []float64{}
		for _, c := range row.Cells {
			v, err := strconv.ParseFloat(c.Value, 64)
			if err != nil {
				return err
			}
			r = append(r, v)
		}
		rows = append(rows, r)
	}

	matrices[mat] = matrix.FromRows(rows)
	return nil
}

func matrixCreate4(mat string, data *godog.Table) error {
	return matrixCreate(4, 4, mat, data)
}

func matrixEquals(m0, m1 string) error {
	mat0 := matrices[m0].(matrix.Mat4)
	mat1 := matrices[m1].(matrix.Mat4)
	if !mat0.Equals(mat1) {
		return fmt.Errorf("Expected %s \n%s to equal %s \n%s",
			m0, mat0, m1, mat1)
	}
	return nil
}

func matrixNotEquals(m0, m1 string) error {
	mat0 := matrices[m0].(matrix.Mat4)
	mat1 := matrices[m1].(matrix.Mat4)
	if mat0.Equals(mat1) {
		return fmt.Errorf("Expected %s \n%s to not equal %s \n%s",
			m0, mat0, m1, mat1)
	}
	return nil
}
