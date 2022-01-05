Feature: Abstract Shapes

Scenario: The default transformation
  Given s ← test_shape()
  Then s.transform = identity_matrix

Scenario: Assigning a transformation
  Given s ← test_shape()
  When set_transform(s, translation(2.0, 3.0, 4.0))
  Then s.transform = translation(2.0, 3.0, 4.0)

Scenario: The default material
  Given s ← test_shape()
  When m ← s.material
  Then m = material()

Scenario: Assigning a material
  Given s ← test_shape()
    And m ← material()
    And m.ambient ← 1.0
  When s.material ← m
  Then s.material = m

Scenario: Intersecting a scaled shape with a ray
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← test_shape()
  When set_transform(s, scaling(2.0, 2.0, 2.0))
    And xs ← intersect(s, r)
  Then s.saved_ray.origin = point(0.0, 0.0, -2.5)
    And s.saved_ray.direction = vector(0.0, 0.0, 0.5)

Scenario: Intersecting a scaled shape with a ray
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← test_shape()
  When set_transform(s, scaling(2.0, 2.0, 2.0))
    And xs ← intersect(s, r)
  Then s.saved_ray.origin = point(0.0, 0.0, -2.5)
    And s.saved_ray.direction = vector(0.0, 0.0, 0.5)

Scenario: Intersecting a translated shape with a ray
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And s ← test_shape()
  When set_transform(s, translation(5.0, 0.0, 0.0))
    And xs ← intersect(s, r)
  Then s.saved_ray.origin = point(-5.0, 0.0, -5.0)
    And s.saved_ray.direction = vector(0.0, 0.0, 1.0)

Scenario: Computing the normal on a translated shape
  Given s ← test_shape()
  When set_transform(s, translation(0.0, 1.0, 0.0))
    And n ← normal_at(s, point(0.0, 1.70711, -0.70711))
  Then n = vector(0.0, 0.70711, -0.70711)

Scenario: Computing the normal on a transformed shape
  Given s ← test_shape()
    And m ← scaling(1.0, 0.5, 1.0) * rotation_z(0.62831853071)
  When set_transform(s, m)
    And n ← normal_at(s, point(0.0, 0.70710678118, -0.70710678118))
  Then n = vector(0.0, 0.97014, -0.24254)
