package matrix

import (
	"fmt"
	"math"

	"github.com/nicholasblaskey/raytracer/tuple"
)

// TODO make all epsilons reference same value (and use same float equals)
const epsilon = 0.0001

func FloatEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

// Column major.
type Matrix interface {
	At(int, int) float64
	Det() float64
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

func Ident4() Mat4 {
	return Mat4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

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

func (a Mat3) Equals(b Mat3) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (a Mat2) Equals(b Mat2) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (a Mat4) Mul4(b Mat4) Mat4 {
	// For performance reasons.
	return Mat4{
		a[0]*b[0] + a[4]*b[1] + a[8]*b[2] + a[12]*b[3],
		a[1]*b[0] + a[5]*b[1] + a[9]*b[2] + a[13]*b[3],
		a[2]*b[0] + a[6]*b[1] + a[10]*b[2] + a[14]*b[3],
		a[3]*b[0] + a[7]*b[1] + a[11]*b[2] + a[15]*b[3],

		a[0]*b[4] + a[4]*b[5] + a[8]*b[6] + a[12]*b[7],
		a[1]*b[4] + a[5]*b[5] + a[9]*b[6] + a[13]*b[7],
		a[2]*b[4] + a[6]*b[5] + a[10]*b[6] + a[14]*b[7],
		a[3]*b[4] + a[7]*b[5] + a[11]*b[6] + a[15]*b[7],

		a[0]*b[8] + a[4]*b[9] + a[8]*b[10] + a[12]*b[11],
		a[1]*b[8] + a[5]*b[9] + a[9]*b[10] + a[13]*b[11],
		a[2]*b[8] + a[6]*b[9] + a[10]*b[10] + a[14]*b[11],
		a[3]*b[8] + a[7]*b[9] + a[11]*b[10] + a[15]*b[11],

		a[0]*b[12] + a[4]*b[13] + a[8]*b[14] + a[12]*b[15],
		a[1]*b[12] + a[5]*b[13] + a[9]*b[14] + a[13]*b[15],
		a[2]*b[12] + a[6]*b[13] + a[10]*b[14] + a[14]*b[15],
		a[3]*b[12] + a[7]*b[13] + a[11]*b[14] + a[15]*b[15],
	}
}

func (a Mat4) Mul4x1(b tuple.Tuple) tuple.Tuple {
	return tuple.Tuple{
		a[0]*b[0] + a[4]*b[1] + a[8]*b[2] + a[12]*b[3],
		a[1]*b[0] + a[5]*b[1] + a[9]*b[2] + a[13]*b[3],
		a[2]*b[0] + a[6]*b[1] + a[10]*b[2] + a[14]*b[3],
		a[3]*b[0] + a[7]*b[1] + a[11]*b[2] + a[15]*b[3],
	}
}

func (m Mat4) T() Mat4 {
	return Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

// TODO revisit some of these with generics released.
// Could be a cool way to learn them
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

func (m Mat3) String() string {
	out := ""
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			out += fmt.Sprintf("%2f ", m.At(y, x))
		}
		out += "\n"
	}
	return out
}

func (m Mat2) String() string {
	out := ""
	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
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

func (m Mat4) Det() float64 {
	return m.Cofactor(0, 0)*m[0] +
		m.Cofactor(0, 1)*m[4] +
		m.Cofactor(0, 2)*m[8] +
		m.Cofactor(0, 3)*m[12]
}

// TOOD optimize this.
func (m Mat4) SubMatrix(x, y int) Mat3 {
	m3 := Mat3{}
	m3I := 0
	for i := 0; i < len(m); i++ {
		if x == i%4 || y == i/4 {
			continue
		}

		m3[m3I] = m[i]
		m3I++
	}

	return m3
}

func (m Mat4) Minor(x, y int) float64 {
	return m.SubMatrix(x, y).Det()
}

func (m Mat4) Cofactor(x, y int) float64 {
	if (x+y)%2 == 0 {
		return m.Minor(x, y)
	}
	return -m.Minor(x, y)
}

func (m Mat3) Det() float64 {
	return m.Cofactor(0, 0)*m[0] +
		m.Cofactor(0, 1)*m[3] +
		m.Cofactor(0, 2)*m[6]
}

func (m Mat3) SubMatrix(x, y int) Mat2 {
	m2 := Mat2{}
	m2I := 0
	for i := 0; i < len(m); i++ {
		if x == i%3 || y == i/3 {
			continue
		}

		m2[m2I] = m[i]
		m2I++
	}

	return m2
}

func (m Mat3) Minor(x, y int) float64 {
	return m.SubMatrix(x, y).Det()
}

func (m Mat3) Cofactor(x, y int) float64 {
	if (x+y)%2 == 0 {
		return m.Minor(x, y)
	}
	return -m.Minor(x, y)
}

func (m Mat2) Det() float64 {
	return m[0]*m[3] - m[1]*m[2]
}
