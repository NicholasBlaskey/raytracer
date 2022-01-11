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

Scenario Outline: A ray strikes a cylinder
  Given cyl ← cylinder()
    And direction ← normalize(<direction>)
    And r ← ray(<origin>, direction)
  When xs ← local_intersect(cyl, r)
  Then xs.count = 2
    And xs[0].t = <t0>
    And xs[1].t = <t1>

  Examples:
    | origin                | direction             |       t0 |      t1 |
    | point(1.0, 0.0, -5.0) | vector(0.0, 0.0, 1.0) |      5.0 |     5.0 |
    | point(0.0, 0.0, -5.0) | vector(0.0, 0.0, 1.0) |      4.0 |     6.0 |
    | point(0.5, 0.0, -5.0) | vector(0.1, 1.0, 1.0) | 6.807982 | 7.08872 |

Scenario Outline: Normal vector on a cylinder
  Given cyl ← cylinder()
  When n ← local_normal_at(cyl, <point>)
  Then n = <normal>

  Examples:
    | point                 | normal                 |
    | point(1.0, 0.0, 0.0)  | vector(1.0, 0.0, 0.0)  |
    | point(0.0, 5.0, -1.0) | vector(0.0, 0.0, -1.0) |
    | point(0.0, -2.0, 1.0) | vector(0.0, 0.0, 1.0)  |
    | point(-1.0, 1.0, 0.0) | vector(-1.0, 0.0, 0.0) |

