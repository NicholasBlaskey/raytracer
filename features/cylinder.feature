Feature: Cylinders

Scenario Outline: A ray misses a cylinder
  Given cyl ← cylinder()
    And direction ← normalize(<direction>)
    And r ← ray(<origin>, direction)
  When xs ← local_intersect(cyl, r)
  Then xs.count = 0

  Examples:
    | origin                | direction             |
    | point(1.0, 0.0, 0.0)  | vector(0.0, 1.0, 0.0) |
    | point(0.0, 0.0, 0.0)  | vector(0.0, 1.0, 0.0) |
    | point(0.0, 0.0, -5.0) | vector(1.0, 1.0, 1.0) |
