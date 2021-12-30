package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/material"

	"github.com/cucumber/godog"
)

var materials map[string]material.Material

func materialBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	materials = make(map[string]material.Material)
	return ctx, nil
}

func materialSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← material\(\)$`,
		wordRegex), createMaterial)

	ctx.Step(fmt.Sprintf(`^%s\.(ambient|diffuse|specular|shininess) = %s$`,
		wordRegex, floatRegex), materialComponentEqual)
	ctx.Step(fmt.Sprintf(`^%s\.color = color\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), materialColorEqual)
}

func createMaterial(m string) {
	materials[m] = material.New()
}

func materialComponentEqual(m, component string, expected float64) error {
	actual := -1.0
	switch component {
	case "ambient":
		actual = materials[m].Ambient
	case "diffuse":
		actual = materials[m].Diffuse
	case "specular":
		actual = materials[m].Specular
	case "shininess":
		actual = materials[m].Shininess
	default:
		panic("SHOULDNT happen!")
	}

	if actual != expected {
		return fmt.Errorf("%s.%s expected %f got %f", m, component, expected, actual)
	}
	return nil
}

func materialColorEqual(m string, r, g, b float64) error {
	actual := fmt.Sprintf("%s.color", m)
	tuples[actual] = materials[m].Color

	return isEqualTuple(actual, r, g, b, 0.0)
}
