package material

import (
	"math"

	"github.com/nicholasblaskey/raytracer/light"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Material struct {
	Color     tuple.Tuple
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
	Pattern   Pattern
}

func New() *Material {
	return &Material{
		Color:     tuple.Color(1.0, 1.0, 1.0),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (m *Material) Lighting(light light.Point, point, eyev, normalv tuple.Tuple,
	inShadow bool) tuple.Tuple {

	col := m.Color
	if m.Pattern != nil {
		col = m.Pattern.At(point)
	}

	effColor := col.ColorMul(light.Intensity)

	lightv := light.Position.Sub(point).Normalize()

	ambient := effColor.Mul(m.Ambient)

	if inShadow {
		return ambient
	}

	lightDotNormal := lightv.Dot(normalv)
	diffuse, specular := tuple.Color(0.0, 0.0, 0.0), tuple.Color(0.0, 0.0, 0.0)
	if lightDotNormal >= 0.0 {
		diffuse = effColor.Mul(m.Diffuse).Mul(lightDotNormal)

		reflectv := lightv.Mul(-1).Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)

		if reflectDotEye > 0.0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Mul(m.Specular).Mul(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
