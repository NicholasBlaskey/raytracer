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
	matrices["identity_matrix"] = matrix.Ident4()

	return ctx, nil
}

func matrixSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^the following %sx%s matrix %s:$`,
		intRegex, intRegex, wordRegex), matrixCreate)
	ctx.Step(fmt.Sprintf(`^the following matrix %s:$`, wordRegex), matrixCreate4)
	ctx.Step(fmt.Sprintf(`^%s\[%s,%s\] = %s$`,
		wordRegex, intRegex, intRegex, floatRegex), matrixElementEqual)

	ctx.Step(fmt.Sprintf(`^%s is the following 4x4 matrix:$`,
		wordRegex), matrixEqual4)
	ctx.Step(fmt.Sprintf(`^%s = %s$`,
		wordRegex, wordRegex), matrixEquals)
	ctx.Step(fmt.Sprintf(`^%s != %s$`,
		wordRegex, wordRegex), matrixNotEquals)

	ctx.Step(fmt.Sprintf(`^%s ← %s \* %s$`,
		wordRegex, wordRegex, wordRegex), matrixMul)
	ctx.Step(fmt.Sprintf(`^%s \* %s is the following 4x4 matrix:$`,
		wordRegex, wordRegex), matrixMulEquals)
	ctx.Step(fmt.Sprintf(`^%s \* %s = %s$`,
		wordRegex, wordRegex, wordRegex), matrixMulEqualsVar)
	ctx.Step(fmt.Sprintf(`%s \* %s = tuple\(%s, %s, %s, %s\)`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		matrixMulTupleEquals)

	ctx.Step(fmt.Sprintf(`^%s ← transpose\(%s\)$`,
		wordRegex, wordRegex), matrixTranspose)
	ctx.Step(fmt.Sprintf(`^transpose\(%s\) is the following matrix:$`,
		wordRegex), matrixTransposeEqual)

	ctx.Step(fmt.Sprintf(`^determinant\(%s\) = %s$`,
		wordRegex, floatRegex), matrixDeterminatEqual)
	ctx.Step(fmt.Sprintf(`^%s ← submatrix\(%s, %s, %s\)$`,
		wordRegex, wordRegex, intRegex, intRegex), matrixSubMatrix)
	ctx.Step(fmt.Sprintf(`^submatrix\(%s, %s, %s\) is the following %sx%s matrix:$`,
		wordRegex, intRegex, intRegex, intRegex, intRegex), matrixSubMatrixIs)
	ctx.Step(fmt.Sprintf(`^minor\(%s, %s, %s\) = %s$`,
		wordRegex, intRegex, intRegex, floatRegex), matrixMinorIs)
	ctx.Step(fmt.Sprintf(`^cofactor\(%s, %s, %s\) = %s$`,
		wordRegex, intRegex, intRegex, floatRegex), matrixCofactorEqual)

	ctx.Step(fmt.Sprintf(`^%s (is|is not) invertible$`,
		wordRegex), matrixCanInv)
	ctx.Step(fmt.Sprintf(`^%s ← inverse\(%s\)$`,
		wordRegex, wordRegex), matrixInv)
	ctx.Step(fmt.Sprintf(`^inverse\(%s\) is the following 4x4 matrix:$`,
		wordRegex), matrixInvEquals)
	ctx.Step(fmt.Sprintf(`^%s \* inverse\(%s\) = %s`,
		wordRegex, wordRegex, wordRegex), matrixInvMulEquals)
}

func matrixElementEqual(mat string, i, j int, expected float64) error {
	if actual := matrices[mat].At(i, j); !compareFloat(actual, expected) {
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

func matrixEqual4(mat string, data *godog.Table) error {
	expected := fmt.Sprintf("expected%s", mat)
	err := matrixCreate(4, 4, expected, data)
	if err != nil {
		return err
	}

	return matrixEquals(mat, expected)
}

func matrixEquals(m0, m1 string) error {
	areEqual := false
	switch mat0 := matrices[m0].(type) {
	case matrix.Mat4:
		areEqual = mat0.Equals(matrices[m1].(matrix.Mat4))
	case matrix.Mat3:
		areEqual = mat0.Equals(matrices[m1].(matrix.Mat3))
	case matrix.Mat2:
		areEqual = mat0.Equals(matrices[m1].(matrix.Mat2))
	}

	if !areEqual {
		return fmt.Errorf("Expected %s \n%s to equal %s \n%s",
			m0, matrices[m0], m1, matrices[m1])
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

func matrixMul(res, m0, m1 string) {
	// Matrix x Matrix case.
	if m, ok := matrices[m1]; ok {
		matrices[res] = matrices[m0].(matrix.Mat4).Mul4(m.(matrix.Mat4))
		return
	}

	// Matrix x tuple case.
	tuples[res] = matrices[m0].(matrix.Mat4).Mul4x1(tuples[m1])
}

func matrixMulEquals(m0, m1 string, data *godog.Table) error {
	expectedM := fmt.Sprintf("expected%s*%s", m0, m1)
	err := matrixCreate4(expectedM, data)
	if err != nil {
		return err
	}
	return matrixMulEqualsVar(m0, m1, expectedM)
}

func matrixMulEqualsVar(m0, m1, expected string) error {
	mat0 := matrices[m0].(matrix.Mat4)

	actual := fmt.Sprintf("actual%s*%s", m0, m1)
	// Matrix * Matrix case
	if mat1, ok := matrices[m1]; ok {
		matrices[actual] = mat0.Mul4(mat1.(matrix.Mat4))

		return matrixEquals(actual, expected)
	}

	// Matrix * tuple case
	a := mat0.Mul4x1(tuples[m1])
	return isEqualTuple(expected, a[0], a[1], a[2], a[3])
}

func matrixMulTupleEquals(m, t string, x, y, z, w float64) error {
	actual := fmt.Sprintf("%s * %s", m, t)
	tuples[actual] = matrices[m].(matrix.Mat4).Mul4x1(tuples[t])

	return isEqualTuple(actual, x, y, z, w)
}

func matrixTranspose(m0, m1 string) {
	matrices[m0] = matrices[m1].(matrix.Mat4).T()
}

func matrixTransposeEqual(mat string, data *godog.Table) error {
	actual := fmt.Sprintf("actual transpose(%s)", mat)
	matrixTranspose(actual, mat)

	expected := fmt.Sprintf("expected transpose(%s)", mat)
	if err := matrixCreate4(expected, data); err != nil {
		return nil
	}
	return matrixEquals(actual, expected)
}

func matrixDeterminatEqual(mat string, expected float64) error {
	actual := matrices[mat].Det()
	if actual != expected {
		return fmt.Errorf("det(%s) expected %f got %f", mat, expected, actual)
	}
	return nil
}

func matrixCofactorEqual(mat string, row, col int, expected float64) error {
	actual := -1.0
	switch m := matrices[mat].(type) {
	case matrix.Mat4:
		actual = m.Cofactor(row, col)
	case matrix.Mat3:
		actual = m.Cofactor(row, col)
	}
	if actual != expected {
		return fmt.Errorf("cofactor(%s, %d, %d) expected %f got %f",
			mat, row, col, expected, actual)
	}
	return nil
}

func matrixSubMatrix(m0, m1 string, row, col int) {
	switch m := matrices[m1].(type) {
	case matrix.Mat4:
		matrices[m0] = m.SubMatrix(row, col)
	case matrix.Mat3:
		matrices[m0] = m.SubMatrix(row, col)
	}
}

func matrixSubMatrixIs(mat string, row, col, n, m int, data *godog.Table) error {
	actual := fmt.Sprintf("actual submatrix(%s, %d, %d)", mat, row, col)
	matrixSubMatrix(actual, mat, row, col)

	expected := fmt.Sprintf("expected submatrix(%s, %d, %d)", mat, row, col)
	if err := matrixCreate(n, m, expected, data); err != nil {
		return nil
	}
	return matrixEquals(actual, expected)
}

func matrixMinorIs(mat string, row, col int, expected float64) error {
	actual := -1.0
	switch m := matrices[mat].(type) {
	case matrix.Mat4:
		actual = m.Minor(row, col)
	case matrix.Mat3:
		actual = m.Minor(row, col)
	}

	if actual != expected {
		return fmt.Errorf("minor(%s, %d, %d) expected %f got %f",
			mat, row, col, expected, actual)
	}
	return nil
}

func matrixCanInv(mat string, isOrIsNot string) error {
	actual := matrices[mat].(matrix.Mat4).HasInv()
	expected := isOrIsNot == "is"
	if actual != expected {
		return fmt.Errorf("%s.HasInv() expected %t got %t", mat,
			expected, actual)
	}
	return nil
}

func matrixInv(m0, m1 string) {
	matrices[m0] = matrices[m1].(matrix.Mat4).Inv()
}

func matrixInvEquals(mat string, data *godog.Table) error {
	expected := fmt.Sprintf("expected Inverse(%s)", mat)
	if err := matrixCreate(4, 4, expected, data); err != nil {
		return err
	}

	actual := fmt.Sprintf("actual Inverse(%s)", mat)
	matrixInv(actual, mat)

	return matrixEquals(actual, expected)
}

func matrixInvMulEquals(m0, m1 string, expected string) error {
	actual := fmt.Sprintf("%s * inverse(%s)", m0, m1)
	matrices[actual] = matrices[m0].(matrix.Mat4).Mul4(
		matrices[m1].(matrix.Mat4).Inv())

	return matrixEquals(actual, expected)
}
