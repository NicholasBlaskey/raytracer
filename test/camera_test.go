package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/camera"

	"github.com/cucumber/godog"
)

var cameras map[string]*camera.Camera
var ints map[string]int
var floats map[string]float64

func cameraBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	cameras = make(map[string]*camera.Camera)
	ints = make(map[string]int)
	floats = make(map[string]float64)

	return ctx, nil
}

func cameraSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← %s$`, wordRegex, intRegex), intCreate)
	ctx.Step(fmt.Sprintf(`^%s ← %s$`, wordRegex, floatRegex), floatCreate)
	ctx.Step(fmt.Sprintf(`^%s ← camera\(%s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex), createCameraFromVars)

	ctx.Step(fmt.Sprintf(`^%s.(hsize|vsize) = %s$`, wordRegex, intRegex),
		cameraIntEqual)
	ctx.Step(fmt.Sprintf(`^%s.field_of_view = %s$`, wordRegex, floatRegex),
		cameraFloatEqual)
}

func intCreate(res string, x int) {
	ints[res] = x
}

func floatCreate(res string, x float64) {
	floats[res] = x
}

func createCameraFromVars(c, hsize, vsize, fov string) {
	cameras[c] = camera.New(ints[hsize], ints[vsize], floats[fov])
}

func cameraIntEqual(c, hSizeOrVSize string, expected int) error {
	actual := cameras[c].HSize
	if hSizeOrVSize == "vsize" {
		actual = cameras[c].VSize
	}

	if actual != expected {
		return fmt.Errorf("%s.%s expected %d got %d", c, hSizeOrVSize, expected, actual)
	}
	return nil
}

func cameraFloatEqual(c string, expected float64) error {
	if !compareFloat(cameras[c].FieldOfView, expected) {
		return fmt.Errorf("%s.field_of_view expected %f got %f", c, expected,
			cameras[c].FieldOfView)
	}
	return nil
}
