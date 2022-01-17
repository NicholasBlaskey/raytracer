package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/obj"
	//"github.com/nicholasblaskey/raytracer/matrix"

	"github.com/cucumber/godog"
)

var fileStrings map[string]string
var objParsers map[string]*obj.Parser

func objBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	fileStrings = make(map[string]string)
	objParsers = make(map[string]*obj.Parser)
	return ctx, nil
}

func objSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← a file containing:$`, wordRegex), createFileString)
	ctx.Step(fmt.Sprintf(`^%s ← parse_obj_file\(%s\)$`, wordRegex, wordRegex),
		parseObjectFile)

	ctx.Step(fmt.Sprintf(`^%s should have ignored %s lines$`, wordRegex, intRegex),
		parserShouldIgnoreLines)
}

func createFileString(f string, contents *godog.DocString) {
	fileStrings[f] = string(contents.Content)
}

func parseObjectFile(p, s string) {
	objParsers[p] = obj.Parse(fileStrings[s])
}

func parserShouldIgnoreLines(p string, expected int) error {
	actual := objParsers[p].LinesIgnored
	if actual != expected {
		return fmt.Errorf("%s.LinesIgnored expected %d got %d", p, expected, actual)
	}
	return nil
}
