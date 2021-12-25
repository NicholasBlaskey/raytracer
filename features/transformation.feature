Feature: Matrix Transformations

Scenario: Multiplying by a translation matrix
  Given transform ← translation(5.0, -3.0, 2.0)
    And p ← point(-3.0, 4.0, 5.0)
   Then transform * p = point(2.0, 1.0, 7.0)

Scenario: Multiplying by the inverse of a translation matrix
  Given transform ← translation(5.0, -3.0, 2.0)
    And inv ← inverse(transform)
    And p ← point(-3.0, 4.0, 5.0)
   Then inv * p = point(-8.0, 7.0, 3.0)

Scenario: Translation does not affect vectors
  Given transform ← translation(5.0, -3.0, 2.0)
    And v ← vector(-3.0, 4.0, 5.0)
   Then transform * v = v
