package matrix

import (
	"fmt"
	"math"
)

// TODO make all epsilons reference same value (and use same float equals)
const epsilon = 0.0001

func FloatEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

// Column major.
type Matrix interface {
	At(int, int) float64
}

func FromRows(rows [][]float64) Matrix {
	if len(rows) == 4 {
		return Mat4{
			rows[0][0], rows[1][0], rows[2][0], rows[3][0],
			rows[0][1], rows[1][1], rows[2][1], rows[3][1],
			rows[0][2], rows[1][2], rows[2][2], rows[3][2],
			rows[0][3], rows[1][3], rows[2][3], rows[3][3],
		}
	}
	if len(rows) == 3 {
		return Mat3{
			rows[0][0], rows[1][0], rows[2][0],
			rows[0][1], rows[1][1], rows[2][1],
			rows[0][2], rows[1][2], rows[2][2],
		}
	}
	return Mat2{rows[0][0], rows[1][0], rows[0][1], rows[1][1]}
}

type Mat4 [16]float64

// TODO profile performance penality of copying on value reciver?
func (m Mat4) At(i, j int) float64 {
	return m[i+j*4]
}

func (a Mat4) Equals(b Mat4) bool {
	// For performance reasons.
	return FloatEquals(a[0], b[0]) &&
		FloatEquals(a[1], b[1]) &&
		FloatEquals(a[2], b[2]) &&
		FloatEquals(a[3], b[3]) &&
		FloatEquals(a[4], b[4]) &&
		FloatEquals(a[5], b[5]) &&
		FloatEquals(a[6], b[6]) &&
		FloatEquals(a[7], b[7]) &&
		FloatEquals(a[8], b[8]) &&
		FloatEquals(a[9], b[9]) &&
		FloatEquals(a[10], b[10]) &&
		FloatEquals(a[11], b[11]) &&
		FloatEquals(a[12], b[12]) &&
		FloatEquals(a[13], b[13]) &&
		FloatEquals(a[14], b[14]) &&
		FloatEquals(a[15], b[15])
}

func (m Mat4) String() string {
	out := ""
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			out += fmt.Sprintf("%2f ", m.At(y, x))
		}
		out += "\n"
	}
	return out
}

type Mat3 [9]float64

func (m Mat3) At(i, j int) float64 {
	return m[i+j*3]
}

type Mat2 [4]float64

func (m Mat2) At(i, j int) float64 {
	return m[i+j*2]
}
