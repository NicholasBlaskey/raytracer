package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/obj"
	"github.com/nicholasblaskey/raytracer/shape"

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

	ctx.Step(fmt.Sprintf(`^%s.vertices\[%s\] = point\(%s, %s, %s\)$`,
		wordRegex, intRegex, floatRegex, floatRegex, floatRegex),
		parserVertexAtEqualTo)

	ctx.Step(fmt.Sprintf(`^%s ← %s.default_group$`, wordRegex, wordRegex),
		assignParserDefaultGroup)
	ctx.Step(fmt.Sprintf(`^%s ← (first|second) child of %s$`,
		wordRegex, wordRegex), assignChildOfGroup)

	ctx.Step(fmt.Sprintf(`^%s.%s = %s.vertices\[%s\]$`,
		wordRegex, wordRegex, wordRegex, intRegex),
		trianglePointEqualToParserVert)
}

func createFileString(f string, contents *godog.DocString) {
	fileStrings[f] = string(contents.Content)
}

func parseObjectFile(p, s string) error {
	parser, err := obj.Parse(fileStrings[s])
	if err != nil {
		return err
	}

	objParsers[p] = parser
	return nil
}

func parserShouldIgnoreLines(p string, expected int) error {
	actual := objParsers[p].LinesIgnored
	if actual != expected {
		return fmt.Errorf("%s.LinesIgnored expected %d got %d", p, expected, actual)
	}
	return nil
}

func parserVertexAtEqualTo(p string, i int, x, y, z float64) error {
	i-- // Tests assume index 1 based scheme
	if i >= len(objParsers[p].Vertices) {
		return fmt.Errorf("%s.vertices[%d] out of bounds (len(%s.vertexes) = %d)",
			p, i, p, len(objParsers[p].Vertices))
	}

	actual := fmt.Sprintf("%s.vertices[%d]", p, i)
	tuples[actual] = objParsers[p].Vertices[i]

	return isEqualTuple(actual, x, y, z, tuples[actual][3])
}

func assignParserDefaultGroup(g, p string) {
	shapes[g] = objParsers[p].DefaultGroup
}

func assignChildOfGroup(res, nth, g string) error {
	i := 0
	switch nth {
	case "first":
		i = 0
	case "second":
		i = 1
	default:
		return fmt.Errorf("Should not happen")
	}

	children := shapes[g].(*shape.Group).Children
	if i >= len(children) {
		return fmt.Errorf("out of bounds wanted %d element of %s.children (len = %d)",
			i, g, len(children))
	}

	shapes[res] = children[i]
	return nil
}

func trianglePointEqualToParserVert(tri, point, parser string, i int) error {
	i--
	if i >= len(objParsers[parser].Vertices) {
		return fmt.Errorf("%s.vertices[%d] out of bounds (len(%s.vertexes) = %d)",
			parser, i, parser, len(objParsers[parser].Vertices))
	}

	expected := fmt.Sprintf("%s.vertices[%d]", parser, i)
	tuples[expected] = objParsers[parser].Vertices[i]
	return trianglePointEqualTo(tri, point, expected)
}
