package main_test

import (
	"context"
	"fmt"

	//"github.com/nicholasblaskey/raytracer/canvas"
	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

var patterns map[string]material.Pattern

func patternBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	patterns = make(map[string]material.Pattern)

	tuples["black"] = tuple.Color(0.0, 0.0, 0.0)
	tuples["white"] = tuple.Color(1.0, 1.0, 1.0)

	return ctx, nil
}

func patternSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← stripe_pattern\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), createStripePattern)
	ctx.Step(fmt.Sprintf(`^%s.(a|b) = %s$`,
		wordRegex, wordRegex), stripePatternColEqual)

	ctx.Step(fmt.Sprintf(`^stripe_at\(%s, point\(%s, %s, %s\)\) = %s$`,
		wordRegex, floatRegex, floatRegex, floatRegex, wordRegex), stripePatternAtEqual)
}

func createStripePattern(p, c1, c2 string) {
	patterns[p] = material.StripePattern(tuples[c1], tuples[c2])
}

func stripePatternColEqual(p, aOrB, expected string) error {
	var actual tuple.Tuple
	if aOrB == "a" {
		actual = patterns[p].(*material.Stripe).Color1
	} else {
		actual = patterns[p].(*material.Stripe).Color2
	}

	return isEqualTuple(expected, actual[0], actual[1], actual[2], actual[3])
}

func stripePatternAtEqual(p string, x, y, z float64, expected string) error {
	actual := patterns[p].At(tuple.Point(x, y, z))
	return isEqualTuple(expected, actual[0], actual[1], actual[2], actual[3])
}
