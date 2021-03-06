package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/material"
	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

var materials map[string]*material.Material
var stringVals map[string]string

func materialBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	materials = make(map[string]*material.Material)
	stringVals = make(map[string]string)

	materials["m"] = material.New() // TODO figure out how to get background properly working
	return ctx, nil
}

func materialSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← material\(\)$`,
		wordRegex), createMaterial)
	ctx.Step(fmt.Sprintf(`^%s\.(transparency|refractive_index|ambient|diffuse|specular|shininess|reflective) ← %s`,
		wordRegex, wordRegex), assignMaterialComponenent)

	ctx.Step(fmt.Sprintf(`^%s ← (true|false)$`, wordRegex),
		assignBooleanString)

	ctx.Step(fmt.Sprintf(`^%s\.(transparency|refractive_index|ambient|diffuse|specular|shininess|reflective) = %s$`,
		wordRegex, floatRegex), materialComponentEqual)
	ctx.Step(fmt.Sprintf(`^%s\.color = color\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), materialColorEqual)

	ctx.Step(fmt.Sprintf(`^%s ← lighting\(%s, %s, %s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex, wordRegex, wordRegex),
		materialLighting)
	ctx.Step(fmt.Sprintf(`^%s ← lighting\(%s, %s, point\(%s, %s, %s\), %s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, floatRegex, floatRegex, floatRegex, wordRegex,
		wordRegex, wordRegex), materialLightingPosLiteral)
	ctx.Step(fmt.Sprintf(`^%s ← lighting\(%s, %s, %s, %s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex, wordRegex, wordRegex, wordRegex),
		materialLightingWithShadowFlag)

	ctx.Step(fmt.Sprintf(`^%s.pattern ← stripe_pattern\(color\(%s, %s, %s\), color\(%s, %s, %s\)\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		assignMaterialPatternStripes)
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
	case "reflective":
		actual = materials[m].Reflective
	case "transparency":
		actual = materials[m].Transparency
	case "refractive_index":
		actual = materials[m].RefractiveIndex
	default:
		panic("SHOULDNT happen!")
	}

	if actual != expected {
		return fmt.Errorf("%s.%s expected %f got %f", m, component, expected, actual)
	}
	return nil
}

func assignMaterialComponenent(m, component string, v float64) {
	switch component {
	case "ambient":
		materials[m].Ambient = v
	case "diffuse":
		materials[m].Diffuse = v
	case "specular":
		materials[m].Specular = v
	case "shininess":
		materials[m].Shininess = v
	case "reflective":
		materials[m].Reflective = v
	case "transparency":
		materials[m].Transparency = v
	case "refractive_index":
		materials[m].RefractiveIndex = v
	default:
		panic("SHOULDNT happen!")
	}
}

func materialColorEqual(m string, r, g, b float64) error {
	actual := fmt.Sprintf("%s.color", m)
	tuples[actual] = materials[m].Color

	return isEqualTuple(actual, r, g, b, 0.0)
}

func assignBooleanString(res, s string) {
	stringVals[res] = s
}

func materialLighting(res, m, light, pos, eyev, normalv string) {
	materialLightingWithShadowFlag(res, m, light, pos, eyev, normalv, "false")
}

func materialLightingPosLiteral(res, m, light string, posX, posY, posZ float64,
	eyev, normalv, inShadow string) {

	pos := "lighting position"
	tuples[pos] = tuple.Point(posX, posY, posZ)
	materialLightingWithShadowFlag(res, m, light, pos, eyev, normalv, inShadow)
}

func materialLightingWithShadowFlag(res, m, light, pos, eyev, normalv, inShadow string) {
	tuples[res] = materials[m].Lighting(
		shape.NewSphere(), // Place holder does nothing
		lights[light],
		tuples[pos],
		tuples[eyev],
		tuples[normalv],
		stringVals[inShadow] == "true",
	)
}

func assignMaterialPatternStripes(m string, r0, g0, b0, r1, g1, b1 float64) {
	materials[m].Pattern = material.StripePattern(tuple.Color(r0, g0, b0),
		tuple.Color(r1, g1, b1))
}
