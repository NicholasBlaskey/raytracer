Feature: Smooth Triangles

Background:
  Given p1 ← point(0.0, 1.0, 0.0)
    And p2 ← point(-1.0, 0.0, 0.0)
    And p3 ← point(1.0, 0.0, 0.0)
    And n1 ← vector(0.0, 1.0, 0.0)
    And n2 ← vector(-1.0, 0.0, 0.0)
    And n3 ← vector(1.0, 0.0, 0.0)
  When tri ← smooth_triangle(p1, p2, p3, n1, n2, n3)

Scenario: Constructing a smooth triangle
  Then tri.p1 = p1
    And tri.p2 = p2
    And tri.p3 = p3
    And tri.n1 = n1
    And tri.n2 = n2
    And tri.n3 = n3

