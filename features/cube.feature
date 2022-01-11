Feature: Cubes

Scenario Outline: A ray intersects a cube
  Given c ← cube()
    And r ← ray(<origin>, <direction>)
  When xs ← local_intersect(c, r)
  Then xs.count = 2
    And xs[0].t = <t1>
    And xs[1].t = <t2>
  Examples:
    |        | origin                 | direction              |   t1 |  t2 |
    | +x     | point(5.0, 0.5, 0.0)   | vector(-1.0, 0.0, 0.0) |  4.0 | 6.0 |
    | -x     | point(-5.0, 0.5, 0.0)  | vector(1.0, 0.0, 0.0)  |  4.0 | 6.0 |
    | +y     | point(0.5, 5.0, 0.0)   | vector(0.0, -1.0, 0.0) |  4.0 | 6.0 |
    | -y     | point(0.5, -5.0, 0.0)  | vector(0.0, 1.0, 0.0)  |  4.0 | 6.0 |
    | +z     | point(0.5, 0.0, 5.0)   | vector(0.0, 0.0, -1.0) |  4.0 | 6.0 |
    | -z     | point(0.5, 0.0, -5.0)  | vector(0.0, 0.0, 1.0)  |  4.0 | 6.0 |
    | inside | point(0.0, 0.5, 0.0)   | vector(0.0, 0.0, 1.0)  | -1.0 | 1.0 |
