Feature: World

Scenario: Creating a world
  Given w ← world()
  Then w contains no objects
    And w has no light source

Scenario: The default world
  Given light ← point_light(point(-10.0, 10.0, -10.0), color(1.0, 1.0, 1.0))
    And s1 ← sphere() with:
      | material.color     | (0.8, 1.0, 0.6)        |
      | material.diffuse   | 0.7                    |
      | material.specular  | 0.2                    |
    And s2 ← sphere() with:
      | transform | scaling(0.5, 0.5, 0.5) |
  When w ← default_world()
  Then w.light = light
    And w contains s1
    And w contains s2

Scenario: Intersect a world with a ray
  Given w ← default_world()
    And r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
  When xs ← intersect_world(w, r)
  Then xs.count = 4
    And xs[0].t = 4.0
    And xs[1].t = 4.5
    And xs[2].t = 5.5
    And xs[3].t = 6.0

Scenario: Shading an intersection
  Given w ← default_world()
    And r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And shape ← the first object in w
    And i ← intersection(4.0, shape)
  When comps ← prepare_computations(i, r)
    And c ← shade_hit(w, comps)
  Then c = color(0.38066, 0.47583, 0.2855)

Scenario: Shading an intersection from the inside
  Given w ← default_world()
    And w.light ← point_light(point(0.0, 0.25, 0.0), color(1.0, 1.0, 1.0))
    And r ← ray(point(0.0, 0.0, 0.0), vector(0.0, 0.0, 1.0))
    And shape ← the second object in w
    And i ← intersection(0.50, shape) 
  When comps ← prepare_computations(i, r)
    And c ← shade_hit(w, comps)
  Then c = color(0.90498, 0.90498, 0.90498)

Scenario: The color when a ray misses
  Given w ← default_world()
    And r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 1.0, 0.0))
  When c ← color_at(w, r)
  Then c = color(0.0, 0.0, 0.0)

Scenario: The color when a ray hits
  Given w ← default_world()
    And r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
  When c ← color_at(w, r)
  Then c = color(0.38066, 0.47583, 0.2855)

Scenario: The color with an intersection behind the ray
  Given w ← default_world()
    And outer ← the first object in w
    And outer.material.ambient ← 1.0
    And inner ← the second object in w
    And inner.material.ambient ← 1.0
    And r ← ray(point(0.0, 0.0, 0.75), vector(0.0, 0.0, -1.0))
  When c ← color_at(w, r)
  Then c = inner.material.color

Scenario: There is no shadow when nothing is collinear with point and light
  Given w ← default_world()
    And p ← point(0.0, 10.0, 0.0)
   Then is_shadowed(w, p) is false

Scenario: The shadow when an object is between the point and the light
  Given w ← default_world()
    And p ← point(10.0, -10.0, 10.0)
   Then is_shadowed(w, p) is true

Scenario: There is no shadow when an object is behind the light
  Given w ← default_world()
    And p ← point(-20.0, 20.0, -20.0)
   Then is_shadowed(w, p) is false

Scenario: There is no shadow when an object is behind the point
  Given w ← default_world()
    And p ← point(-2.0, 2.0, -2.0)
   Then is_shadowed(w, p) is false

Scenario: shade_hit() is given an intersection in shadow
  Given w ← world()
    And w.light ← point_light(point(0.0, 0.0, -10.0), color(1.0, 1.0, 1.0))
    And s1 ← sphere()
    And s1 is added to w
    And s2 ← sphere() with:
      | transform | translation(0, 0, 10) |
    And s2 is added to w
    And r ← ray(point(0.0, 0.0, 5.0), vector(0.0, 0.0, 1.0))
    And i ← intersection(4.0, s2)
  When comps ← prepare_computations(i, r)
    And c ← shade_hit(w, comps)
  Then c = color(0.1, 0.1, 0.1)

Scenario: The reflected color for a nonreflective material
  Given w ← default_world()
    And r ← ray(point(0.0, 0.0, 0.0), vector(0.0, 0.0, 1.0))
    And shape ← the second object in w
    And shape.material.ambient ← 1.0
    And i ← intersection(1.0, shape)
  When comps ← prepare_computations(i, r)
    And color ← reflected_color(w, comps)
  Then color = color(0.0, 0.0, 0.0)

Scenario: The reflected color for a reflective material
  Given w ← default_world()
    And shape ← plane() with:                 
      | material.reflective | 0.5                   |
      | transform           | translation(0, -1, 0) |   
    And shape is added to w
    And r ← ray(point(0.0, 0.0, -3.0), vector(0.0, -0.70710678118, 0.70710678118))
    And i ← intersection(1.41421356237, shape)
  When comps ← prepare_computations(i, r)
    And color ← reflected_color(w, comps)
  Then color = color(0.190331, 0.237913, 0.142748)
