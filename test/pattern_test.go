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
	ctx.Step(fmt.Sprintf(`^%s ← test_pattern\(\)$`, wordRegex), createTestPattern)
	ctx.Step(fmt.Sprintf(`^%s ← gradient_pattern\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), createGradientPattern)
	ctx.Step(fmt.Sprintf(`^%s ← ring_pattern\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), createRingPattern)
	ctx.Step(fmt.Sprintf(`^%s ← checkers_pattern\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), createCheckerPattern)

	ctx.Step(fmt.Sprintf(`^%s.(a|b) = %s$`,
		wordRegex, wordRegex), stripePatternColEqual)
	ctx.Step(fmt.Sprintf(`^%s ← (?:pattern_at_shape|stripe_at_object)\(%s, %s, point\(%s, %s, %s\)\)$`,
		wordRegex, wordRegex, wordRegex, floatRegex, floatRegex, floatRegex), stripeAtObject)
	ctx.Step(fmt.Sprintf(`^pattern_at\(%s, point\(%s, %s, %s\)\) = color\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		patternAtEqualColor)

	ctx.Step(fmt.Sprintf(`^(?:pattern_at|stripe_at)\(%s, point\(%s, %s, %s\)\) = %s$`,
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

func createTestPattern(p string) {
	patterns[p] = &TestPattern{matrix.Ident4()}
}

func patternAtEqualColor(p string, x, y, z, r, g, b float64) error {
	actual := fmt.Sprintf("pattern_at(%s, point(%f, %f, %f))", p, x, y, z)
	tuples[actual] = patterns[p].At(tuple.Point(x, y, z))

	// Take A componenet of actual in case we need to refactor that to be 1s for colors.
	return isEqualTuple(actual, r, g, b, tuples[actual][3])
}

func createGradientPattern(p, c0, c1 string) {
	patterns[p] = material.GradientPattern(tuples[c0], tuples[c1])
}

func createRingPattern(p, c0, c1 string) {
	patterns[p] = material.RingPattern(tuples[c0], tuples[c1])
}

func createCheckerPattern(p, c0, c1 string) {
	patterns[p] = material.CheckerPattern(tuples[c0], tuples[c1])
}

type TestPattern struct {
	Transform matrix.Mat4
}

func (s *TestPattern) GetTransform() matrix.Mat4 {
	return s.Transform
}

func (s *TestPattern) SetTransform(m matrix.Mat4) {
	s.Transform = m
}

func (s *TestPattern) At(p tuple.Tuple) tuple.Tuple {
	return tuple.Color(p[0], p[1], p[2])
}

func (s *TestPattern) AtObject(obj material.Object, worldPoint tuple.Tuple) tuple.Tuple {
	return s.At(material.WorldToPattern(s, obj, worldPoint))
}
