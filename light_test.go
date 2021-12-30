package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/light"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

var lights map[string]light.Point

func lightBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	lights = make(map[string]light.Point)
	return ctx, nil
}

func lightSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← point_light\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), createLight)
	ctx.Step(fmt.Sprintf(`^%s ← point_light\(point\(%s, %s, %s\), color\(%s, %s, %s\)\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex,
		floatRegex, floatRegex, floatRegex), createLightFromLiteral)

	ctx.Step(fmt.Sprintf(`^%s.(position|intensity) = %s$`,
		wordRegex, wordRegex), lightComponentEqual)
}

func createLight(l, pos, intensity string) {
	lights[l] = light.NewPointLight(tuples[pos], tuples[intensity])
}

func createLightFromLiteral(l string, x, y, z, r, g, b float64) {
	lights[l] = light.NewPointLight(tuple.Point(x, y, z), tuple.Color(r, g, b))
}

func lightComponentEqual(l, posOrIntensity, expected string) error {
	actual := lights[l].Intensity
	if posOrIntensity == "position" {
		actual = lights[l].Position
	}

	return isEqualTuple(expected, actual[0], actual[1], actual[2], actual[3])
}
