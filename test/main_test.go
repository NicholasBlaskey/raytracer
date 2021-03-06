package main_test

import (
	"context"
	"math"
	"testing"

	"github.com/cucumber/godog"
)

const (
	epsilon    = 0.0001
	wordRegex  = `([_A-Za-z0-9^\s]+)` // TODO change this to no spaces? Why did we have no spaces?
	intRegex   = `(\d+)`
	floatRegex = `(\-*\d+\.\d+)`
)

func compareFloat(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: initializeScenario,

		Options: &godog.Options{
			Format:   "pretty", //"progress", // pretty
			Paths:    []string{"../features/csg.feature"},
			TestingT: t, // Testing instance that will run subtests.
			// Stops on the first failure
			//StopOnFailure: true,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func initializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tuplesBefore(ctx, sc)
		canvasBefore(ctx, sc)
		matrixBefore(ctx, sc)
		transformBefore(ctx, sc)
		rayBefore(ctx, sc)
		sphereBefore(ctx, sc)
		intersectionBefore(ctx, sc)
		lightBefore(ctx, sc)
		materialBefore(ctx, sc)
		worldBefore(ctx, sc)
		cameraBefore(ctx, sc)
		shapeBefore(ctx, sc)
		planeBefore(ctx, sc)
		patternBefore(ctx, sc)
		cubeBefore(ctx, sc)
		cylinderBefore(ctx, sc)
		groupBefore(ctx, sc)
		triangleBefore(ctx, sc)
		objBefore(ctx, sc)
		smoothTriangleBefore(ctx, sc)
		csgBefore(ctx, sc)

		return ctx, nil
	})

	tuplesSteps(ctx)
	canvasSteps(ctx)
	matrixSteps(ctx)
	transformSteps(ctx)
	raySteps(ctx)
	sphereSteps(ctx)
	intersectionSteps(ctx)
	lightSteps(ctx)
	materialSteps(ctx)
	worldSteps(ctx)
	cameraSteps(ctx)
	shapeSteps(ctx)
	planeSteps(ctx)
	patternSteps(ctx)
	cubeSteps(ctx)
	cylinderSteps(ctx)
	groupSteps(ctx)
	triangleSteps(ctx)
	objSteps(ctx)
	smoothTriangleSteps(ctx)
	csgSteps(ctx)
}
