package main_test

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nicholasblaskey/raytracer/matrix"
	//"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

/*
var mat4s map[string]matrix.Mat4
var mat3s map[string]matrix.Mat3
var mat2s map[string]matrix.Mat2
*/
var matrices map[string]matrix.Matrix

func matrixBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	matrices = make(map[string]matrix.Matrix)
	/*
		mat4s = make(map[string]matrix.Mat4)
		mat3s = make(map[string]matrix.Mat3)
		mat2s = make(map[string]matrix.Mat2)
	*/

	return ctx, nil
}

func matrixSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^the following %sx%s matrix %s:$`,
		intRegex, intRegex, wordRegex), matrixCreate)
	ctx.Step(fmt.Sprintf(`^%s\[%s,%s\] = %s$`,
		wordRegex, intRegex, intRegex, floatRegex), matrixElementEqual)
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
