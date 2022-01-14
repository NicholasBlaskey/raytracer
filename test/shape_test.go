package main_test

import (
	"context"
	"fmt"
	//"strconv"
	//"strings"

	"github.com/nicholasblaskey/raytracer/intersection"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/tuple"

	/*
		"github.com/nicholasblaskey/raytracer/material"
		"github.com/nicholasblaskey/raytracer/matrix"
		"github.com/nicholasblaskey/raytracer/shape"
		"github.com/nicholasblaskey/raytracer/tuple"
	*/

	"github.com/cucumber/godog"
)

var shapes map[string]intersection.Intersectable

//var intersections map[string][]*intersection.Intersection

func shapeBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	shapes = make(map[string]intersection.Intersectable)

	return ctx, nil
}

func shapeSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← test_shape\(\)$`, wordRegex), createTestShape)

	ctx.Step(fmt.Sprintf(`^%s.saved_ray.(origin|direction) = (point|vector)\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), savedRayEqual)
}

func createTestShape(s string) {
	shapes[s] = testShapeNew()
}

func savedRayEqual(s, origOrDir, pointOrVec string, x, y, z float64) error {
	return rayComponentEqualVecOrPoint("saved_ray", origOrDir, pointOrVec, x, y, z)
}

type testShape struct {
	Transform matrix.Mat4
	Material  *material.Material
	Parent    intersection.Intersectable
}

func testShapeNew() *testShape {
	return &testShape{Transform: matrix.Ident4(), Material: material.New()}
}

func (s *testShape) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *testShape) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func (s *testShape) GetMaterial() *material.Material {
	return s.Material
}

func (s *testShape) SetMaterial(m *material.Material) {
	s.Material = m
}

func (s *testShape) GetParent() intersection.Intersectable {
	return s.Parent
}

func (s *testShape) SetParent(p intersection.Intersectable) {
	s.Parent = p
}

func (s *testShape) NormalAt(t tuple.Tuple) tuple.Tuple {
	return intersection.NormalAt(s, t,
		func(objectPoint tuple.Tuple) tuple.Tuple {
			return tuple.Vector(objectPoint[0], objectPoint[1], objectPoint[2]).Normalize()
		})
}

func (s *testShape) Intersections(origR ray.Ray) []*intersection.Intersection {
	return intersection.Intersections(s, origR, func(r ray.Ray) []*intersection.Intersection {
		rays["saved_ray"] = r
		return nil
	})
}
