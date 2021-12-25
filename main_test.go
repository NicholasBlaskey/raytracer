package main_test

import (
	"context"
	"math"
	"testing"

	"github.com/cucumber/godog"
)

const (
	epsilon    = 0.00001
	wordRegex  = `([_A-Za-z0-9^\s]+)`
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
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
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

		return ctx, nil
	})

	tuplesSteps(ctx)
	canvasSteps(ctx)
	matrixSteps(ctx)
	transformSteps(ctx)
}
