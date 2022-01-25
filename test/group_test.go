package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/shape"

	"github.com/cucumber/godog"
)

func groupBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func groupSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← group\(\)$`, wordRegex), createGroup)

	ctx.Step(fmt.Sprintf(`^add_child\(%s, %s\)$`,
		wordRegex, wordRegex), groupAddChild)
	ctx.Step(fmt.Sprintf(`^%s is not empty$`, wordRegex), groupIsNotEmpty)
	ctx.Step(fmt.Sprintf(`^%s includes %s$`,
		wordRegex, wordRegex), groupIncludesChild)
	ctx.Step(fmt.Sprintf(`^%s.parent = %s$`,
		wordRegex, wordRegex), shapeParentEqualTo)
}

func createGroup(g string) {
	shapes[g] = shape.NewGroup()
}

func groupAddChild(g, s string) {
	group := shapes[g].(*shape.Group)
	group.AddChild(shapes[s])
}

func groupIsNotEmpty(g string) error {
	if group := shapes[g].(*shape.Group); group.Children == nil {
		return fmt.Errorf("%s expected to not be empty got %v instead",
			g, group.Children)
	}
	return nil
}

func groupIncludesChild(g, s string) error {
	group := shapes[g].(*shape.Group)
	found := false
	for _, c := range group.Children {
		if c == shapes[s] {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("%s.children (%v) did not include %s", g, group.Children, s)
	}
	return nil
}

func shapeParentEqualTo(s, g string) error {
	if shapes[s].GetParent() != shapes[g] {
		return fmt.Errorf("%s.parent expected %+v got %+v", s,
			shapes[s].GetParent(), shapes[g])
	}

	return nil
}
