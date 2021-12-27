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
