package main_test

import (
	"context"
	"fmt"
	"io/ioutil"

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
	ctx.Step(fmt.Sprintf(`^%s ← the file "%s\.obj"$`, wordRegex, wordRegex),
		fileIntoString)

	ctx.Step(fmt.Sprintf(`^%s ← a file containing:$`, wordRegex), createFileString)
	ctx.Step(fmt.Sprintf(`^%s ← parse_obj_file\(%s\)$`, wordRegex, wordRegex),
		parseObjectFile)

	ctx.Step(fmt.Sprintf(`^%s should have ignored %s lines$`, wordRegex, intRegex),
		parserShouldIgnoreLines)

	ctx.Step(fmt.Sprintf(`^%s.vertices\[%s\] = point\(%s, %s, %s\)$`,
		wordRegex, intRegex, floatRegex, floatRegex, floatRegex),
		parserVertexAtEqualTo)
	ctx.Step(fmt.Sprintf(`^%s.normals\[%s\] = vector\(%s, %s, %s\)$`,
		wordRegex, intRegex, floatRegex, floatRegex, floatRegex),
		parserNormalAtEqualTo)

	ctx.Step(fmt.Sprintf(`^%s ← %s.default_group$`, wordRegex, wordRegex),
		assignParserDefaultGroup)
	ctx.Step(fmt.Sprintf(`^%s ← (first|second|third) child of %s$`,
		wordRegex, wordRegex), assignChildOfGroup)
	ctx.Step(fmt.Sprintf(`^%s ← "%s" from %s$`,
		wordRegex, wordRegex, wordRegex), getNamedGroupFromParser)

	ctx.Step(fmt.Sprintf(`^%s.%s = %s.vertices\[%s\]$`,
		wordRegex, wordRegex, wordRegex, intRegex),
		trianglePointEqualToParserVert)
	ctx.Step(fmt.Sprintf(`^%s.%s = %s.normals\[%s\]$`,
		wordRegex, wordRegex, wordRegex, intRegex),
		smoothTriangleNormalEqualTo)

	ctx.Step(fmt.Sprintf(`^%s ← obj_to_group\(%s\)$`,
		wordRegex, wordRegex), objToGroup)
	ctx.Step(fmt.Sprintf(`^%s includes "%s" from %s$`,
		wordRegex, wordRegex, wordRegex), groupIncludesParserGroup)
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

func parserNormalAtEqualTo(p string, i int, x, y, z float64) error {
	i--
	if i >= len(objParsers[p].Normals) {
		return fmt.Errorf("%s.normals[%d] out of bounds (len(%s.normals) = %d)",
			p, i, p, len(objParsers[p].Normals))
	}

	actual := fmt.Sprintf("%s.normals[%d]", p, i)
	tuples[actual] = objParsers[p].Normals[i]

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
	case "third":
		i = 2
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

func smoothTriangleNormalEqualTo(tri, point, parser string, i int) error {
	i--
	if i >= len(objParsers[parser].Normals) {
		return fmt.Errorf("%s.normals[%d] out of bounds (len(%s.normals) = %d)",
			parser, i, parser, len(objParsers[parser].Normals))
	}

	expected := fmt.Sprintf("%s.normals[%d]", parser, i)
	tuples[expected] = objParsers[parser].Normals[i]

	t := shapes[tri].(*shape.SmoothTriangle)
	actual := t.N0
	if point == "n2" {
		actual = t.N1
	} else if point == "n3" {
		actual = t.N2
	}

	return isEqualTuple(expected, actual[0], actual[1], actual[2], actual[3])
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

func fileIntoString(f, fileName string) error {
	b, err := ioutil.ReadFile(fileName + ".obj")
	if err != nil {
		return err
	}

	fileStrings[f] = string(b)
	return nil
}

func getNamedGroupFromParser(res, groupName, parser string) error {
	group := objParsers[parser].Groups[groupName]
	if group == nil {
		return fmt.Errorf("expected %s to have group %s, but did not", parser, groupName)
	}

	shapes[res] = group
	return nil
}

func objToGroup(res, parser string) {
	shapes[res] = objParsers[parser].ToGroup()
}

func groupIncludesParserGroup(group, nameOfGroup, parser string) error {
	parserGroup := objParsers[parser].Groups[nameOfGroup]

	found := false
	for _, c := range shapes[group].(*shape.Group).Children {
		if c == parserGroup {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("expected %s to have %s from %s", group, nameOfGroup, parser)
	}
	return nil
}
