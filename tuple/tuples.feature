Feature: Tuples, Vectors, and Points

Scenario: A tuple with w=1.0 is a point
  Given a ← tuple(4.3, -4.2, 3.1, 1.0)
  Then a.x = 4.3
    And a.y = -4.2
    And a.z = 3.1
    And a.w = 1.0
    And a is a point
    And a is not a vector

Scenario: A tuple with w=0 is a vector
  Given a ← tuple(4.3, -4.2, 3.1, 0.0)
  Then a.x = 4.3
    And a.y = -4.2
    And a.z = 3.1
    And a.w = 0.0
    And a is not a point
    And a is a vector

Scenario: point() creates tuples with w=1
  Given p ← point(4.0, -4.0, 3.0)
  Then p = tuple(4.0, -4.0, 3.0, 1.0)

Scenario: vector() creates tuples with w=0
  Given v ← vector(4.0, -4.0, 3.0)
  Then v = tuple(4.0, -4.0, 3.0, 0.0)

Scenario: Adding two tuples
  Given a1 ← tuple(3.0, -2.0, 5.0, 1.0)
    And a2 ← tuple(-2.0, 3.0, 1.0, 0.0)
   Then a1 + a2 = tuple(1.0, 1.0, 6.0, 1.0)

Scenario: Subtracting two points
  Given p1 ← point(3.0, 2.0, 1.0)
    And p2 ← point(5.0, 6.0, 7.0)
  Then p1 - p2 = vector(-2.0, -4.0, -6.0)

Scenario: Subtracting a vector from a point
  Given p ← point(3.0, 2.0, 1.0)
    And v ← vector(5.0, 6.0, 7.0)
  Then p - v = point(-2.0, -4.0, -6.0)

Scenario: Subtracting two vectors
  Given v1 ← vector(3.0, 2.0, 1.0)
    And v2 ← vector(5.0, 6.0, 7.0)
  Then v1 - v2 = vector(-2.0, -4.0, -6.0)

Scenario: Subtracting a vector from the zero vector
  Given zero ← vector(0.0, 0.0, 0.0)
    And v ← vector(1.0, -2.0, 3.0)
  Then zero - v = vector(-1.0, 2.0, -3.0)

Scenario: Negating a tuple
  Given a ← tuple(1.0, -2.0, 3.0, -4.0)
  Then -a = tuple(-1.0, 2.0, -3.0, 4.0)
