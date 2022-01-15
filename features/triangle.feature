Feature: Triangles

Scenario: Constructing a triangle
  Given p1 ← point(0.0, 1.0, 0.0)
    And p2 ← point(-1.0, 0.0, 0.0)
    And p3 ← point(1.0, 0.0, 0.0)
    And t ← triangle(p1, p2, p3)
  Then t.p1 = p1
    And t.p2 = p2
    And t.p3 = p3
    And t.e1 = vector(-1.0, -1.0, 0.0)
    And t.e2 = vector(1.0, -1.0, 0.0)
    And t.normal = vector(0.0, 0.0, -1.0)
