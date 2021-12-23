package matrix

import (
//"fmt"
)

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

type Mat3 [9]float64

func (m Mat3) At(i, j int) float64 {
	return m[i+j*3]
}

type Mat2 [4]float64

func (m Mat2) At(i, j int) float64 {
	return m[i+j*2]
}
