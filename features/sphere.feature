Feature: Spheres

Scenario: A ray intersects a sphere at two points
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0] = 4.0
    And xs[1] = 6.0

Scenario: A ray intersects a sphere at a tangent
  Given r ← ray(point(0.0, 1.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0] = 5.0
    And xs[1] = 5.0

Scenario: A ray misses a sphere
  Given r ← ray(point(0.0, 2.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 0

Scenario: A ray originates inside a sphere
  Given r ← ray(point(0.0, 0.0, 0.0), vector(0.0, 0.0, 1.0))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0] = -1.0
    And xs[1] = 1.0

Scenario: A sphere is behind a ray
  Given r ← ray(point(0.0, 0.0, 5.0), vector(0.0, 0.0, 1.0))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0] = -6.0
    And xs[1] = -4.0

Scenario: Intersect sets the object on the intersection
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0].object = s
    And xs[1].object = s

Scenario: A sphere's default transformation
  Given s ← sphere()
  Then s.transform = identity_matrix

Scenario: Changing a sphere's transformation
  Given s ← sphere()
    And t ← translation(2.0, 3.0, 4.0)
  When set_transform(s, t)
  Then s.transform = t

Scenario: Intersecting a scaled sphere with a ray
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← sphere()
  When set_transform(s, scaling(2.0, 2.0, 2.0))
    And xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0].t = 3.0
    And xs[1].t = 7.0

Scenario: Intersecting a translated sphere with a ray
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← sphere()
  When set_transform(s, translation(5.0, 0.0, 0.0))
    And xs ← intersect(s, r)
  Then xs.count = 0

Scenario: The normal on a sphere at a point on the x axis
  Given s ← sphere()
  When n ← normal_at(s, point(1.0, 0.0, 0.0))
  Then n = vector(1.0, 0.0, 0.0)

Scenario: The normal on a sphere at a point on the y axis
  Given s ← sphere()
  When n ← normal_at(s, point(0.0, 1.0, 0.0))
  Then n = vector(0.0, 1.0, 0.0)

Scenario: The normal on a sphere at a point on the z axis
  Given s ← sphere()
  When n ← normal_at(s, point(0.0, 0.0, 1.0))
  Then n = vector(0.0, 0.0, 1.0)

Scenario: The normal on a sphere at a nonaxial point
  Given s ← sphere()
  When n ← normal_at(s, point(0.57735026919, 0.57735026919, 0.57735026919))
  Then n = vector(0.57735026919, 0.57735026919, 0.57735026919)

Scenario: The normal is a normalized vector
  Given s ← sphere()
  When n ← normal_at(s, point(0.57735026919, 0.57735026919, 0.57735026919))
  Then n = normalize(n)