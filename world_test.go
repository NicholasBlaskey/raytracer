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

var worlds map[string]*world.World

func worldBefore(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	worlds = make(map[string]*world.World)
	return ctx, nil
}

func worldSteps(ctx *godog.ScenarioContext) {
	ctx.Step(fmt.Sprintf(`^%s ← world\(\)$`, wordRegex), createWorld)
	ctx.Step(fmt.Sprintf(`^%s ← default_world\(\)$`, wordRegex), defaultWorld)

	ctx.Step(fmt.Sprintf(`^%s\.light = %s$`, wordRegex, wordRegex), worldLightEqual)
	ctx.Step(fmt.Sprintf(`^%s contains ([A-Za-z0-9]+)$`, wordRegex), worldContains) // TODO fix this to use word regex

	ctx.Step(fmt.Sprintf(`^%s contains no objects$`, wordRegex), worldHasNoObjects)
	ctx.Step(fmt.Sprintf(`^%s has no light source$`, wordRegex), worldHasNoLightSource)
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

func worldLightEqual(w, l string) error {
	if worlds[w].Light == nil {
		return fmt.Errorf("%s.light expected %+v got nil", w, lights[l])
	}
	if *worlds[w].Light != lights[l] {
		return fmt.Errorf("%s.light expected %+v got %+v", w, worlds[w].Light, lights[l])
	}
	return nil
}

func worldContains(w, s string) error {
	found := false
	for _, obj := range worlds[w].Objects {
		fmt.Println(obj.(*shape.Sphere).Transform == spheres[s].Transform,
			*(obj.(*shape.Sphere).Material) == *(spheres[s].Material))

		if obj.(*shape.Sphere).Transform == spheres[s].Transform &&
			*(obj.(*shape.Sphere).Material) == *(spheres[s].Material) {
			found = true
			break
		}
	}

	fmt.Println("EXPECTED", s)
	fmt.Println(spheres[s].Material)
	fmt.Println(spheres[s].Transform)

	fmt.Println("ACTUAL 0")
	fmt.Println(worlds[w].Objects[0].(*shape.Sphere).Material)
	fmt.Println(worlds[w].Objects[0].(*shape.Sphere).Transform)

	fmt.Println("ACTUAL 1")
	fmt.Println(worlds[w].Objects[1].(*shape.Sphere).Material)
	fmt.Println(worlds[w].Objects[1].(*shape.Sphere).Transform)

	/*
		fmt.Println(spheres[s].Material)
		fmt.Println(worlds[w].Objects[0].(*shape.Sphere).Material)
		fmt.Println(worlds[w].Objects[1].(*shape.Sphere).Material)
		fmt.Println(
			*(spheres[s].Material) == *(worlds[w].Objects[1].(*shape.Sphere).Material))
	*/

	if !found {
		return fmt.Errorf("%s does not contain object %+v has only %+v", w,
			spheres[s], worlds[w].Objects)
	}
	return nil
}
