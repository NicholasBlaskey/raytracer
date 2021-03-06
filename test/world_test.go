package main_test

import (
	"context"
	"fmt"

	"github.com/nicholasblaskey/raytracer/light"
	"github.com/nicholasblaskey/raytracer/matrix"
	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"
	"github.com/nicholasblaskey/raytracer/world"

	"github.com/cucumber/godog"
)

const maxDepth = 5

var worlds map[string]*world.World

func worldBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	worlds = make(map[string]*world.World)
	return ctx, nil
}

func worldSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← world\(\)$`, wordRegex), createWorld)
	ctx.Step(fmt.Sprintf(`^%s ← default_world\(\)$`, wordRegex), defaultWorld)
	ctx.Step(fmt.Sprintf(`^%s ← intersect_world\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), intersectWorld)
	ctx.Step(fmt.Sprintf(`^%s.light ← point_light\(point\(%s, %s, %s\), color\(%s, %s, %s\)\)$`,
		wordRegex, floatRegex, floatRegex, floatRegex, floatRegex, floatRegex,
		floatRegex), worldPointLightLiteral)

	ctx.Step(fmt.Sprintf(`^%s\.light = %s$`, wordRegex, wordRegex), worldLightEqual)
	ctx.Step(fmt.Sprintf(`^%s contains ([A-Za-z0-9]+)$`, wordRegex), worldContains) // TODO fix this to use word regex
	ctx.Step(fmt.Sprintf(`^%s is added to %s$`, wordRegex, wordRegex), addToWorld)

	ctx.Step(fmt.Sprintf(`^%s contains no objects$`, wordRegex), worldHasNoObjects)
	ctx.Step(fmt.Sprintf(`^%s has no light source$`, wordRegex), worldHasNoLightSource)

	ctx.Step(fmt.Sprintf(`^%s ← the (first|second) object in %s`,
		wordRegex, wordRegex), theNthObjectFromWorld)
	ctx.Step(fmt.Sprintf(`^%s ← shade_hit\(%s, %s\)`,
		wordRegex, wordRegex, wordRegex), shadeHit)
	ctx.Step(fmt.Sprintf(`^%s ← shade_hit\(%s, %s, %s\)`,
		wordRegex, wordRegex, wordRegex, intRegex), shadeHitWithRemaining)

	ctx.Step(fmt.Sprintf(`^%s.material.ambient ← %s$`,
		wordRegex, floatRegex), setObjectAmbientTo)
	ctx.Step(fmt.Sprintf(`^%s = %s.material.color$`,
		wordRegex, wordRegex), colorEqualToMaterialColor)

	ctx.Step(fmt.Sprintf(`^%s ← color_at\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), worldColorAt)

	ctx.Step(fmt.Sprintf(`^is_shadowed\(%s, %s\) is (true|false)$`,
		wordRegex, wordRegex), isShadowed)
	ctx.Step(fmt.Sprintf(`^%s ← reflected_color\(%s, %s\)$`,
		wordRegex, wordRegex, wordRegex), worldReflectedColor)
	ctx.Step(fmt.Sprintf(`^%s ← reflected_color\(%s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, intRegex), worldReflectedColorDepthLimit)
	ctx.Step(fmt.Sprintf(`^%s ← refracted_color\(%s, %s, %s\)$`,
		wordRegex, wordRegex, wordRegex, intRegex), worldRefractedColor)

	ctx.Step(fmt.Sprintf(`^color_at\(%s, %s\) should terminate successfully$`,
		wordRegex, wordRegex), colorAtShouldHault)
}

func createWorld(w string) {
	worlds[w] = world.New()
}

func worldHasNoObjects(w string) error {
	if len(worlds[w].Objects) != 0 {
		return fmt.Errorf("expected %s to have no objects got %v", w, worlds[w].Objects)
	}
	return nil
}

func worldHasNoLightSource(w string) error {
	if worlds[w].Light != nil {
		return fmt.Errorf("expected %s to have no light source got %v instead",
			w, worlds[w].Light)
	}
	return nil
}

func defaultWorld(w string) {
	createWorld(w)

	l := light.NewPointLight(tuple.Point(-10.0, 10.0, -10.0), tuple.Color(1.0, 1.0, 1.0))
	worlds[w].Light = &l

	s0 := shape.NewSphere()
	s0.Material.Color = tuple.Color(0.8, 1.0, 0.6)
	s0.Material.Diffuse = 0.7
	s0.Material.Specular = 0.2

	s1 := shape.NewSphere()
	s1.Transform = matrix.Scale(0.5, 0.5, 0.5)

	worlds[w].Objects = append(worlds[w].Objects, s0, s1)
}

func worldPointLightLiteral(w string, x, y, z, r, g, b float64) {
	l := light.NewPointLight(tuple.Point(x, y, z),
		tuple.Color(r, g, b))
	worlds[w].Light = &l
}

func worldLightEqual(w, l string) error {
	if worlds[w].Light == nil {
		return fmt.Errorf("%s.light expected %+v got nil", w, lights[l])
	}
	if *worlds[w].Light != lights[l] {
		return fmt.Errorf("%s.light expected %+v got %+v", w, worlds[w].Light, lights[l])
	}
	return nil
}

func addToWorld(s, w string) {
	worlds[w].Objects = append(worlds[w].Objects, shapes[s])
}

func worldContains(w, s string) error {
	found := false
	for _, obj := range worlds[w].Objects {
		if obj.(*shape.Sphere).Transform == shapes[s].GetTransform() &&
			*(obj.(*shape.Sphere).Material) == *(shapes[s].GetMaterial()) {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%s does not contain object %+v has only %+v", w,
			shapes[s], worlds[w].Objects)
	}
	return nil
}

func intersectWorld(res, w, r string) {
	intersections[res] = worlds[w].Intersect(rays[r])
}

func theNthObjectFromWorld(res, nth, w string) {
	switch nth {
	case "first":
		shapes[res] = worlds[w].Objects[0]
	case "second":
		shapes[res] = worlds[w].Objects[1]
	default:
		panic("Unspported nth object from world :" + nth)
	}

}

func shadeHit(res, w, comps string) {
	tuples[res] = worlds[w].ShadeHit(computations[comps], maxDepth)
}

func shadeHitWithRemaining(res, w, comps string, remaining int) {
	tuples[res] = worlds[w].ShadeHit(computations[comps], remaining)
}

func colorAtShouldHault(w, r string) {
	worlds[w].ColorAt(rays[r], 5)
}

func worldColorAt(res, w, r string) {
	tuples[res] = worlds[w].ColorAt(rays[r], maxDepth)
}

func worldReflectedColor(res, w, comps string) {
	tuples[res] = worlds[w].ReflectedColor(computations[comps], maxDepth)
}

func worldRefractedColor(res, w, comps string, remaining int) {
	tuples[res] = worlds[w].RefractedColor(computations[comps], remaining)
}

func worldReflectedColorDepthLimit(res, w, comps string, limit int) {
	tuples[res] = worlds[w].ReflectedColor(computations[comps], limit)
}

func setObjectAmbientTo(obj string, v float64) {
	shapes[obj].GetMaterial().Ambient = v
}

func colorEqualToMaterialColor(c, obj string) error {
	actual := shapes[obj].GetMaterial().Color

	return isEqualTuple(c, actual[0], actual[1], actual[2], actual[3])
}

func isShadowed(w, p, expectedString string) error {
	expected := expectedString == "true"
	if actual := worlds[w].IsShadowed(tuples[p]); actual != expected {
		return fmt.Errorf("is_shadowed(%s, %s) expected %t got %t",
			w, p, expected, actual)
	}
	return nil
}
