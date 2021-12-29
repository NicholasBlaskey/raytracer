Feature: Rays

Scenario: Creating and querying a ray
  Given origin ← point(1.0, 2.0, 3.0)
    And direction ← vector(4.0, 5.0, 6.0)
  When r ← ray(origin, direction)
  Then r.origin = origin
    And r.direction = direction

Scenario: Computing a point from a distance
  Given r ← ray(point(2.0, 3.0, 4.0), vector(1.0, 0.0, 0.0))
  Then position(r, 0.0) = point(2.0, 3.0, 4.0)
    And position(r, 1.0) = point(3.0, 3.0, 4.0)
    And position(r, -1.0) = point(1.0, 3.0, 4.0)
    And position(r, 2.5) = point(4.5, 3.0, 4.0)


Scenario: Translating a ray
  Given r ← ray(point(1.0, 2.0, 3.0), vector(0.0, 1.0, 0.0))
    And m ← translation(3.0, 4.0, 5.0)
  When r2 ← transform(r, m)
  Then r2.origin = point(4.0, 6.0, 8.0)
    And r2.direction = vector(0.0, 1.0, 0.0)

Scenario: Scaling a ray
  Given r ← ray(point(1.0, 2.0, 3.0), vector(0.0, 1.0, 0.0))
    And m ← scaling(2.0, 3.0, 4.0)
  When r2 ← transform(r, m)
  Then r2.origin = point(2.0, 6.0, 12.0)
    And r2.direction = vector(0.0, 3.0, 0.0)
