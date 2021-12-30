package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

var spheres map[string]*shape.Sphere
var intersections map[string][]*intersection.Intersection

func sphereBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	spheres = make(map[string]*shape.Sphere)
	intersections = make(map[string][]*intersection.Intersection)
	return ctx, nil
}

func sphereSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← sphere\(\)$`, wordRegex), createSphere)
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
	ctx.Step(fmt.Sprintf(`^set_transform\(%s, (scaling|translation)\(%s, %s, %s\)\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), sphereSetTransformLiteral)

	ctx.Step(fmt.Sprintf(`^%s ← normal_at\(%s, point\(%s, %s, %s\)\)$`,
		wordRegex, wordRegex, floatRegex, floatRegex, floatRegex), sphereNormalAt)
}

func createSphere(s string) {
	spheres[s] = shape.NewSphere()
}

func intersectSphere(intersection, s, r string) {
	intersections[intersection] = spheres[s].Intersections(rays[r])
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

	if intersections[i][index].Obj != spheres[expected] {
		return fmt.Errorf("%s[%d] expected %v got %v", i, index,
			spheres[expected], intersections[i][index].Obj)
	}
	return nil
}

func sphereSetTransform(s, t string) {
	spheres[s].Transform = matrices[t].(matrix.Mat4)
}

func sphereSetTransformLiteral(s, translationType string, x, y, z float64) {
	switch translationType {
	case "translation":
		spheres[s].Transform = matrix.Translate(x, y, z)
	case "scaling":
		spheres[s].Transform = matrix.Scale(x, y, z)
	}

}

func sphereTransformEquals(s, expected string) error {
	actual := fmt.Sprintf("%s.transform", s)
	matrices[actual] = spheres[s].Transform

	return matrixEquals(actual, expected)
}

func sphereNormalAt(n, s string, x, y, z float64) {
	tuples[n] = spheres[s].NormalAt(tuple.Point(x, y, z))
}
