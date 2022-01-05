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
}

func createTestShape(s string) {
	shapes[s] = testShapeNew()
}

type testShape struct {
	Transform matrix.Mat4
	Material  *material.Material
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

func (s *testShape) NormalAt(t tuple.Tuple) tuple.Tuple {
	return t
}

func (s *testShape) Intersections(r ray.Ray) []*intersection.Intersection {
	return nil
}
