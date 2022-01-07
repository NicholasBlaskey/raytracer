package main_test

import (
	"context"
	"fmt"

	//"github.com/nicholasblaskey/raytracer/ray"
	"github.com/nicholasblaskey/raytracer/intersection"

	"github.com/cucumber/godog"
)

var intersectionObjects map[string]*intersection.Intersection
var computations map[string]*intersection.Computations

func intersectionBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	intersectionObjects = make(map[string]*intersection.Intersection)
	computations = make(map[string]*intersection.Computations)
	return ctx, nil
}

func intersectionSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← intersection\(%s, %s\)$`,
		wordRegex, floatRegex, wordRegex), intersectionCreate)
	// TODO figure out varadic cucumber steps.
	ctx.Step(fmt.Sprintf(`^%s ← intersections\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), aggregateIntersections)
	ctx.Step(fmt.Sprintf(`^%s ← intersections\(%s, %s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, wordRegex, wordRegex),
		aggregateIntersections4)

	ctx.Step(fmt.Sprintf(`^%s ← hit\(%s\)$`, wordRegex, wordRegex),
		intersectionCreateHit)

	ctx.Step(fmt.Sprintf(`^%s\.t = %s$`, wordRegex, floatRegex),
		intersectionTimeEqual)
	ctx.Step(fmt.Sprintf(`^%s.object = %s$`, wordRegex, wordRegex),
		intersectionObjectEqual)

	ctx.Step(fmt.Sprintf(`^%s is nothing$`, wordRegex),
		intersectionIsNothing)

	ctx.Step(fmt.Sprintf(`^%s ← prepare_computations\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), intersectionsPrepareComputations)
	ctx.Step(fmt.Sprintf(`^%s.t = %s\.t$`,
		wordRegex, wordRegex), computationsTimeEquals)
	ctx.Step(fmt.Sprintf(`^%s.inside = (true|false)`, wordRegex),
		computationInsideEquals)
	ctx.Step(fmt.Sprintf(`^%s.object = %s.object$`,
		wordRegex, wordRegex), computationsObjectEquals)
	ctx.Step(fmt.Sprintf(`^%s.(point|eyev|normalv|reflectv) = (point|vector)\(%s, %s, %s\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex), computationsTupleEquals)

	ctx.Step(fmt.Sprintf(`^%s.over_point.(x|y|z) < -EPSILON/2$`, wordRegex),
		overPointLessThanEpsilonOver2)
	ctx.Step(fmt.Sprintf(`^%s.point.(x|y|z) > %s.over_point.(x|y|z)$`,
		wordRegex, wordRegex), pointBiggerThanOverPoint)

}

func intersectionCreate(i string, t float64, obj string) {
	intersectionObjects[i] = intersection.New(t, shapes[obj])
}

func intersectionTimeEqual(i string, t float64) error {
	if actual := intersectionObjects[i].T; actual != t {
		return fmt.Errorf("%s.t expected %f got %f", i, t, actual)
	}
	return nil
}

func intersectionObjectEqual(i, obj string) error {
	if actual := intersectionObjects[i].Obj; actual != shapes[obj] {
		return fmt.Errorf("%s.object expected %v got %v", i, obj, actual)
	}
	return nil
}

func aggregateIntersections(res, i0, i1 string) {
	intersections[res] = intersection.Aggregate(intersectionObjects[i0],
		intersectionObjects[i1])
}

func aggregateIntersections4(res, i0, i1, i2, i3 string) {
	intersections[res] = intersection.Aggregate(intersectionObjects[i0],
		intersectionObjects[i1], intersectionObjects[i2], intersectionObjects[i3])
}

func intersectionCreateHit(hit, i string) {
	intersectionObjects[hit] = intersection.Hit(intersections[i])
}

func intersectionIsNothing(i string) error {
	if intersectionObjects[i] != nil {
		return fmt.Errorf("Expected %s to be null got %v", i, intersectionObjects[i])
	}
	return nil
}

func intersectionEquals(i0, i1 string) error {
	if intersectionObjects[i0] != intersectionObjects[i1] {
		return fmt.Errorf("expected %s (%v) to equal %s (%v)",
			i0, intersectionObjects[i0], i1, intersectionObjects[i1])
	}
	return nil
}

func intersectionsPrepareComputations(res, i, r string) {
	computations[res] = intersectionObjects[i].PrepareComputations(rays[r])
}

func computationsTimeEquals(comp, i string) error {
	if computations[comp].T != intersectionObjects[i].T {
		return fmt.Errorf("%s.t expected %f got %f", comp, intersectionObjects[i].T,
			computations[comp].T)
	}
	return nil
}

func computationsObjectEquals(comp, i string) error {
	fmt.Println(computations[comp], "!", intersectionObjects[i], i)
	if computations[comp].Obj != intersectionObjects[i].Obj {
		return fmt.Errorf("%s.object expected %+v got %+v", comp,
			intersectionObjects[i].Obj, computations[comp].Obj)
	}
	return nil
}

func computationsTupleEquals(comp, component, vectorOrPoint string, x, y, z float64) error {
	actual := fmt.Sprintf("%s.%s", comp, component)
	switch component {
	case "point":
		tuples[actual] = computations[comp].Point
	case "eyev":
		tuples[actual] = computations[comp].Eyev
	case "normalv":
		tuples[actual] = computations[comp].Normalv
	case "reflectv":
		tuples[actual] = computations[comp].Reflectv
	}

	w := 1.0
	if vectorOrPoint == "vector" {
		w = 0.0
	}
	return isEqualTuple(actual, x, y, z, w)
}

func computationInsideEquals(comp string, expected string) error {
	if computations[comp].Inside != (expected == "true") {
		return fmt.Errorf("%s.inside expected %s got %t", comp, expected,
			computations[comp].Inside)
	}
	return nil
}

func overPointLessThanEpsilonOver2(comp, xyz string) error {
	val := computations[comp].OverPoint[0]
	if xyz == "y" {
		val = computations[comp].OverPoint[1]
	} else {
		val = computations[comp].OverPoint[2]
	}

	if val >= -intersection.EPSILON/2.0 {
		return fmt.Errorf("%s.over_point.%s (%f) is not < (%f) -EPSILON/2",
			comp, xyz, val, -intersection.EPSILON/2.0)
	}
	return nil
}

func pointBiggerThanOverPoint(comp, xyz string) error {
	over := computations[comp].OverPoint[0]
	p := computations[comp].Point[0]
	if xyz == "y" {
		p = computations[comp].Point[1]
		over = computations[comp].OverPoint[1]
	} else {
		p = computations[comp].Point[2]
		over = computations[comp].OverPoint[2]
	}

	if p <= over {
		return fmt.Errorf("%s.point.%s (%f) is not < (%f) %s.over_point.%s",
			comp, xyz, p, over, comp, xyz)
	}
	return nil
}
