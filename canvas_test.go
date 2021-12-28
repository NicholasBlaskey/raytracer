package main_test

import (
	"context"
	"fmt"
	"strings"

	"github.com/nicholasblaskey/raytracer/canvas"
	"github.com/nicholasblaskey/raytracer/tuple"

	"github.com/cucumber/godog"
)

var canvases map[string]*canvas.Canvas
var ppms map[string]string

func canvasBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	canvases = make(map[string]*canvas.Canvas)
	ppms = make(map[string]string)
	return ctx, nil
}

func canvasSteps(ctx *godog.ScenarioContext) {
	// Canvas
	ctx.Step(fmt.Sprintf(`^%s ← canvas\(%s, %s\)$`,
		wordRegex, intRegex, intRegex), createCanvas)
	ctx.Step(fmt.Sprintf(`^%s\.(width|height) = %s$`,
		wordRegex, intRegex), checkCanvasWidthOrHeight)
	ctx.Step(fmt.Sprintf(`^every pixel of %s is color\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), canvasEveryPixelIs)
	ctx.Step(fmt.Sprintf(`^write_pixel\(%s, %s, %s, %s\)$`,
		wordRegex, intRegex, intRegex, wordRegex), canvasWritePixel)
	ctx.Step(fmt.Sprintf(`^pixel_at\(%s, %s, %s\) = %s$`,
		wordRegex, intRegex, intRegex, wordRegex), canvasAssertPixel)
	ctx.Step(fmt.Sprintf(`^%s ← canvas_to_ppm\(%s\)$`,
		wordRegex, wordRegex), canvasToPPM)
	ctx.Step(fmt.Sprintf(`^lines %s-%s of %s are$`,
		intRegex, intRegex, wordRegex), ppmLinesAre)
	ctx.Step(fmt.Sprintf(`^every pixel of %s is set to color\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), canvasEveryPixelSetTo)
	ctx.Step(fmt.Sprintf(`^%s ends with a newline character$`,
		wordRegex), ppmEndsWithNewLine)
}

func createCanvas(canv string, w, h int) {
	canvases[canv] = canvas.New(w, h)
}

func checkCanvasWidthOrHeight(canv, widthOrHeight string, v int) error {
	gotten := canvases[canv].Height
	if widthOrHeight == "width" {
		gotten = canvases[canv].Width
	}

	if gotten != v {
		return fmt.Errorf("canvas %s expected %s to be %d got %d",
			canv, widthOrHeight, v, gotten)
	}
	return nil
}

func canvasEveryPixelIs(canv string, x, y, z float64) error {
	c := canvases[canv]
	for i := 0; i < len(c.Pixels); i += 3 {
		if c.Pixels[i] != x || c.Pixels[i+1] != y || c.Pixels[i+2] != z {
			return fmt.Errorf("canvas %s expected color(%f, %f, %f) color(%f, %f, %f)",
				canv, c.Pixels[i], c.Pixels[i+1], c.Pixels[i+2], x, y, z)
		}
	}
	return nil
}

func canvasWritePixel(canv string, x, y int, color string) {
	canvases[canv].WritePixel(tuples[color], x, y)
}

func canvasAssertPixel(canv string, x, y int, color string) error {
	gotten := canvases[canv].ReadPixel(x, y)
	return isEqualTuple(color, gotten[0], gotten[1], gotten[2], gotten[3])
}

func canvasToPPM(ppm string, canv string) {
	ppms[ppm] = canvases[canv].ToPPM()
}

func ppmLinesAre(from, to int, ppm string, lines string) error {
	ppmSplit := strings.Split(ppms[ppm], "\n")
	linesSplit := strings.Split(lines, "\n")
	for i := from; i <= to; i++ {
		actual := ppmSplit[i-1]

		//fmt.Println("ACTUAL LINESPSLIT", ppmSplit, linesSplit)
		//fmt.Println(i, actual, linesSplit[i-from])
		if strings.Trim(actual, " ") != strings.Trim(linesSplit[i-from], " ") {
			return fmt.Errorf("For ppm file %s expected line %d to be \n%s\n got \n%s\n",
				ppm, i, linesSplit[i-from], actual)
		}
	}
	return nil
}

func canvasEveryPixelSetTo(canv string, r, g, b float64) {
	c := canvases[canv]
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			c.WritePixel(tuple.Color(r, g, b), x, y)
		}
	}
}

func ppmEndsWithNewLine(ppm string) error {
	if char := ppms[ppm][len(ppms[ppm])-1]; char != '\n' {
		return fmt.Errorf("PPM ends with %s instead of new line", string(char))
	}
	return nil
}
