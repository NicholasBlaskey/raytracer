Feature: Cones

Scenario Outline: Intersecting a cone with a ray
  Given shape ← cone()
    And direction ← normalize(<direction>)
    And r ← ray(<origin>, direction)
  When xs ← local_intersect(shape, r)
  Then xs.count = 2
    And xs[0].t = <t0>
    And xs[1].t = <t1>

  Examples:
    | origin                | direction               |      t0 |       t1 |
    | point(0.0, 0.0, -5.0) | vector(0.0, 0.0, 1.0)   |     5.0 |      5.0 |
    | point(0.0, 0.0, -5.0) | vector(1.0, 1.0, 1.0)   | 8.66025 |  8.66025 |
    | point(1.0, 1.0, -5.0) | vector(-0.5, -1.0, 1.0) | 4.55006 | 49.44994 |

Scenario: Intersecting a cone with a ray parallel to one of its halves
  Given shape ← cone()
    And direction ← normalize(vector(0.0, 1.0, 1.0))
    And r ← ray(point(0.0, 0.0, -1.0), direction)
  When xs ← local_intersect(shape, r)
  Then xs.count = 1
    And xs[0].t = 0.35355
