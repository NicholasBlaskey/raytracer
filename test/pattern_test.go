package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/matrix"
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

	ctx.Step(fmt.Sprintf(`^%s ← stripe_at_object\(%s, %s, point\(%s, %s, %s\)\)$`,
		wordRegex, wordRegex, wordRegex, floatRegex, floatRegex, floatRegex), stripeAtObject)
	ctx.Step(fmt.Sprintf(`^stripe_at\(%s, point\(%s, %s, %s\)\) = %s$`,
		wordRegex, floatRegex, floatRegex, floatRegex, wordRegex), stripePatternAtEqual)

	ctx.Step(fmt.Sprintf(`^set_pattern_transform\(%s, scaling\(%s, %s, %s\)\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), setPatternTransformScale)
	ctx.Step(fmt.Sprintf(`^set_pattern_transform\(%s, translation\(%s, %s, %s\)\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), setPatternTransformTranslate)

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

func stripeAtObject(res, p, obj string, x, y, z float64) {
	tuples[res] = patterns[p].AtObject(shapes[obj], tuple.Point(x, y, z))
}

func stripePatternAtEqual(p string, x, y, z float64, expected string) error {
	actual := patterns[p].At(tuple.Point(x, y, z))
	return isEqualTuple(expected, actual[0], actual[1], actual[2], actual[3])
}

func setPatternTransformScale(p string, x, y, z float64) {
	patterns[p].SetTransform(matrix.Scale(x, y, z))
}

func setPatternTransformTranslate(p string, x, y, z float64) {
	patterns[p].SetTransform(matrix.Translate(x, y, z))
}
