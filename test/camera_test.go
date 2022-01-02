package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/camera"
	"github.com/nicholasblaskey/raytracer/matrix"

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
	ctx.Step(fmt.Sprintf(`^%s ← camera\(%s, %s, %s\)$`,
		wordRegex, intRegex, intRegex, floatRegex), createCameraFromLiteral)

	ctx.Step(fmt.Sprintf(`^%s ← ray_for_pixel\(%s, %s, %s\)$`,
		wordRegex, wordRegex, intRegex, intRegex), cameraRayForPixel)
	ctx.Step(fmt.Sprintf(`^%s ← render\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), cameraRender)

	ctx.Step(fmt.Sprintf(`^%s.transform ← view_transform\(%s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex), cameraTransformView)
	ctx.Step(fmt.Sprintf(`^%s.transform ← rotation_y\(%s\) \* translation\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex),
		cameraTransformToRotYMulTranslate)

	ctx.Step(fmt.Sprintf(`^%s.(hsize|vsize) = %s$`, wordRegex, intRegex),
		cameraIntEqual)
	ctx.Step(fmt.Sprintf(`^%s.(field_of_view|pixel_size) = %s$`, wordRegex, floatRegex),
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

func createCameraFromLiteral(c string, hsize, vsize int, fov float64) {
	cameras[c] = camera.New(hsize, vsize, fov)
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

func cameraFloatEqual(c, fovOrPixelSize string, expected float64) error {
	actual := cameras[c].FieldOfView
	if fovOrPixelSize == "pixel_size" {
		actual = cameras[c].PixelSize
	}

	if !compareFloat(actual, expected) {
		return fmt.Errorf("%s.%s expected %f got %f", c, fovOrPixelSize,
			expected, actual)
	}
	return nil
}

func cameraTransformToRotYMulTranslate(c string, theta, x, y, z float64) {
	cameras[c].Transform = matrix.RotateY(theta).Mul4(matrix.Translate(x, y, z))
}

func cameraTransformView(c, from, to, up string) {
	cameras[c].Transform = matrix.View(tuples[from], tuples[to], tuples[up])
}

func cameraRayForPixel(r, c string, x, y int) {
	rays[r] = cameras[c].RayForPixel(x, y)
}

func cameraRender(img, c, w string) {
	canvases[img] = cameras[c].Render(worlds[w])
}
