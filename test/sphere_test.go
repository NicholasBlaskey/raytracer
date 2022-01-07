package main_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

//var spheres map[string]*shape.Sphere
var intersections map[string][]*intersection.Intersection

func sphereBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	//	spheres = make(map[string]*shape.Sphere)
	intersections = make(map[string][]*intersection.Intersection)
	return ctx, nil
}

func sphereSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← sphere\(\)$`, wordRegex), createSphere)
	ctx.Step(fmt.Sprintf(`^%s ← (sphere|plane)\(\) with:$`, wordRegex), createShapeWith)
	ctx.Step(fmt.Sprintf(`^%s ← intersect\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), intersectSphere)

	ctx.Step(fmt.Sprintf(`^%s.count = %s`, wordRegex, intRegex),
		intersectCountEqual)
	ctx.Step(fmt.Sprintf(`^%s\[%s\](?:.t|) = %s`, wordRegex, intRegex, floatRegex),
		intersectValueEqual)
	ctx.Step(fmt.Sprintf(`^%s\[%s\].object = %s`, wordRegex, intRegex, wordRegex),
		intersectObjectEqual)

	ctx.Step(fmt.Sprintf(`^set_transform\(%s, %s\)$`, wordRegex, wordRegex),
		sphereSetTransform)

	ctx.Step(fmt.Sprintf(`^%s.transform = %s$`, wordRegex, wordRegex),
		sphereTransformEquals)
	ctx.Step(fmt.Sprintf(`^%s.transform = translation\(%s, %s, %s\)$`, wordRegex,
		floatRegex, floatRegex, floatRegex), sphereTransformEqualsTranslate)
	ctx.Step(fmt.Sprintf(`^set_transform\(%s, (scaling|translation)\(%s, %s, %s\)\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), sphereSetTransformLiteral)

	ctx.Step(fmt.Sprintf(`^%s ← normal_at\(%s, point\(%s, %s, %s\)\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex), sphereNormalAt)

	ctx.Step(fmt.Sprintf(`^%s ← %s.material$`, wordRegex, wordRegex),
		getSphereMaterial)
	ctx.Step(fmt.Sprintf(`^%s = material\(\)$`, wordRegex),
		materialIsDefault)
	ctx.Step(fmt.Sprintf(`^%s.material ← %s$`, wordRegex, wordRegex),
		assignSphereMaterial)
	ctx.Step(fmt.Sprintf(`^%s.material = %s$`, wordRegex, wordRegex),
		sphereMaterialEqual)
}

func createSphere(s string) {
	shapes[s] = shape.NewSphere()
}

func intersectSphere(intersection, s, r string) {
	intersections[intersection] = shapes[s].Intersections(rays[r])
}

func intersectCountEqual(i string, expected int) error {
	actual := len(intersections[i])
	if actual != expected {
		return fmt.Errorf("%s.count expected %d got %d", i, expected, actual)
	}
	return nil
}

func intersectValueEqual(i string, index int, expected float64) error {
	if index >= len(intersections[i]) {
		return fmt.Errorf("Tried to get out of bounds index %d of intersection %v",
			index, len(intersections[i]))
	}

	if intersections[i][index].T != expected {
		return fmt.Errorf("%s[%d] expected %f got %f", i, index,
			expected, intersections[i][index].T)
	}
	return nil
}

func intersectObjectEqual(i string, index int, expected string) error {
	if index >= len(intersections[i]) {
		return fmt.Errorf("Tried to get out of bounds index %d of intersection %v",
			index, len(intersections[i]))
	}

	if intersections[i][index].Obj != shapes[expected] {
		return fmt.Errorf("%s[%d] expected %v got %v", i, index,
			shapes[expected], intersections[i][index].Obj)
	}
	return nil
}

func sphereSetTransform(s, t string) {
	shapes[s].SetTransform(matrices[t].(matrix.Mat4))
}

func sphereSetTransformLiteral(s, translationType string, x, y, z float64) {
	switch translationType {
	case "translation":
		shapes[s].SetTransform(matrix.Translate(x, y, z))
	case "scaling":
		shapes[s].SetTransform(matrix.Scale(x, y, z))
	}

}

func sphereTransformEquals(s, expected string) error {
	actual := fmt.Sprintf("%s.transform", s)

	if c, ok := cameras[s]; ok { // Camera case
		matrices[actual] = c.Transform
	} else if p, ok := patterns[s]; ok { // Pattern case
		matrices[actual] = p.GetTransform()
	} else { // Sphere case
		matrices[actual] = shapes[s].GetTransform()
	}

	return matrixEquals(actual, expected)
}

func sphereTransformEqualsTranslate(s string, x, y, z float64) error {
	actual := fmt.Sprintf("actual %s.transform", s)
	expected := fmt.Sprintf("expected %s.transform", s)

	if p, ok := patterns[s]; ok {
		matrices[actual] = p.GetTransform()
	} else {
		matrices[actual] = shapes[s].GetTransform()
	}
	matrices[expected] = matrix.Translate(x, y, z)

	return matrixEquals(actual, expected)
}

func sphereNormalAt(n, s string, x, y, z float64) {
	tuples[n] = shapes[s].NormalAt(tuple.Point(x, y, z))
}

func getSphereMaterial(m, s string) {
	materials[m] = shapes[s].GetMaterial()
}

func materialIsDefault(m string) error {
	isEqual := *materials[m] == *material.New()
	if !isEqual {
		return fmt.Errorf("%s is not equal to the default material", m)
	}
	return nil
}

func assignSphereMaterial(s, m string) {
	shapes[s].SetMaterial(materials[m])
}

func sphereMaterialEqual(s, m string) error {
	if shapes[s].GetMaterial() != materials[m] {
		return fmt.Errorf("%s.material expected %v got %v", s,
			shapes[s].GetMaterial(), materials[m])
	}
	return nil
}

func parseFloatList(list string) ([]float64, error) {
	// Remove ( )
	list = strings.Split(list, "(")[1]
	list = strings.ReplaceAll(list, ")", "")

	floats := []float64{}
	for _, v := range strings.Split(list, ",") {
		vFloat, err := strconv.ParseFloat(strings.Trim(v, " "), 64)
		if err != nil {
			return nil, err
		}

		floats = append(floats, vFloat)
	}
	return floats, nil
}

func createShapeWith(s, shapeType string, data *godog.Table) error {
	var sph intersection.Intersectable
	if shapeType == "sphere" {
		sph = shape.NewSphere()
	} else if shapeType == "plane" {
		sph = shape.NewPlane()
	}

	for _, row := range data.Rows {
		k, v := row.Cells[0].Value, row.Cells[1].Value
		switch k {
		case "material.color":
			col, err := parseFloatList(v)
			if err != nil {
				panic(err)
			}
			sph.GetMaterial().Color = tuple.Color(col[0], col[1], col[2])
		case "material.diffuse":
			diff, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}

			sph.GetMaterial().Diffuse = diff
		case "material.specular":
			spec, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}

			sph.GetMaterial().Specular = spec
		case "material.reflective":
			r, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}

			sph.GetMaterial().Reflective = r
		case "transform":
			vals, err := parseFloatList(v)
			if err != nil {
				return err
			}

			if strings.Contains(v, "scaling") {
				sph.SetTransform(matrix.Scale(vals[0], vals[1], vals[2]))
			} else if strings.Contains(v, "translation") {
				sph.SetTransform(matrix.Translate(vals[0], vals[1], vals[2]))
			} else if false { // TODO add more matrices here

			} else {
				return fmt.Errorf("Unexpected transform type of %s", v)
			}

		default:
			panic("Unexpected sphere with type of key=" + k + " value=" + v)
		}
	}

	shapes[s] = sph

	return nil
}
