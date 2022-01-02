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

Scenario: A scaling matrix applied to a point
  Given transform ← scaling(2.0, 3.0, 4.0)
    And p ← point(-4.0, 6.0, 8.0)
   Then transform * p = point(-8.0, 18.0, 32.0)

Scenario: A scaling matrix applied to a vector
  Given transform ← scaling(2.0, 3.0, 4.0)
    And v ← vector(-4.0, 6.0, 8.0)
   Then transform * v = vector(-8.0, 18.0, 32.0)

Scenario: Multiplying by the inverse of a scaling matrix
  Given transform ← scaling(2.0, 3.0, 4.0)
    And inv ← inverse(transform)
    And v ← vector(-4.0, 6.0, 8.0)
   Then inv * v = vector(-2.0, 2.0, 2.0)

Scenario: Reflection is scaling by a negative value
  Given transform ← scaling(-1.0, 1.0, 1.0)
    And p ← point(2.0, 3.0, 4.0)
   Then transform * p = point(-2.0, 3.0, 4.0)


Scenario: Rotating a point around the x axis
  Given p ← point(0.0, 1.0, 0.0)                       
                                  # π / 4          
    And half_quarter ← rotation_x(0.78539816339) 
                                 # π / 2
    And full_quarter ← rotation_x(1.57079632679)
  Then half_quarter * p = point(0.0, 0.70710678118, 0.70710678118)
    And full_quarter * p = point(0.0, 0.0, 1.0)

Scenario: The inverse of an x-rotation rotates in the opposite direction
  Given p ← point(0.0, 1.0, 0.0)
                                 # π / 4
    And half_quarter ← rotation_x(0.78539816339)
    And inv ← inverse(half_quarter)
  Then inv * p = point(0.0, 0.70710678118, -0.70710678118)

Scenario: Rotating a point around the y axis
  Given p ← point(0.0, 0.0, 1.0)
                                 # π / 4
    And half_quarter ← rotation_y(0.78539816339)
                                 # π / 2
    And full_quarter ← rotation_y(1.57079632679)
  Then half_quarter * p = point(0.70710678118, 0.0, 0.70710678118)
    And full_quarter * p = point(1.0, 0.0, 0.0)

Scenario: Rotating a point around the z axis
  Given p ← point(0.0, 1.0, 0.0)
                                 # π / 4
    And half_quarter ← rotation_z(0.78539816339)
                                 # π / 2
    And full_quarter ← rotation_z(1.57079632679)
  Then half_quarter * p = point(-0.70710678118, 0.70710678118, 0.0)
    And full_quarter * p = point(-1.0, 0.0, 0.0)

Scenario: A shearing transformation moves x in proportion to y
  Given transform ← shearing(1.0, 0.0, 0.0, 0.0, 0.0, 0.0)
    And p ← point(2.0, 3.0, 4.0)
  Then transform * p = point(5.0, 3.0, 4.0)

Scenario: A shearing transformation moves x in proportion to z
  Given transform ← shearing(0.0, 1.0, 0.0, 0.0, 0.0, 0.0)
    And p ← point(2.0, 3.0, 4.0)
  Then transform * p = point(6.0, 3.0, 4.0)

Scenario: A shearing transformation moves y in proportion to x
  Given transform ← shearing(0.0, 0.0, 1.0, 0.0, 0.0, 0.0)
    And p ← point(2.0, 3.0, 4.0)
  Then transform * p = point(2.0, 5.0, 4.0)

Scenario: A shearing transformation moves y in proportion to z
  Given transform ← shearing(0.0, 0.0, 0.0, 1.0, 0.0, 0.0)
    And p ← point(2.0, 3.0, 4.0)
  Then transform * p = point(2.0, 7.0, 4.0)

Scenario: A shearing transformation moves z in proportion to x
  Given transform ← shearing(0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
    And p ← point(2.0, 3.0, 4.0)
  Then transform * p = point(2.0, 3.0, 6.0)

Scenario: A shearing transformation moves z in proportion to y
  Given transform ← shearing(0.0, 0.0, 0.0, 0.0, 0.0, 1.0)
    And p ← point(2.0, 3.0, 4.0)
  Then transform * p = point(2.0, 3.0, 7.0)

Scenario: Individual transformations are applied in sequence
  Given p ← point(1.0, 0.0, 1.0)
                      # π / 2
    And A ← rotation_x(1.57079632679)
    And B ← scaling(5.0, 5.0, 5.0)
    And C ← translation(10.0, 5.0, 7.0)
  # apply rotation first
  When p2 ← A * p
  Then p2 = point(1.0, -1.0, 0.0)
  # then apply scaling
  When p3 ← B * p2
  Then p3 = point(5.0, -5.0, 0.0)
  # then apply translation
  When p4 ← C * p3
  Then p4 = point(15.0, 0.0, 7.0)

Scenario: Chained transformations must be applied in reverse order
  Given p ← point(1.0, 0.0, 1.0)
                      # π / 2
    And A ← rotation_x(1.57079632679)
    And B ← scaling(5.0, 5.0, 5.0)
    And C ← translation(10.0, 5.0, 7.0)
  When T ← C * B * A
  Then T * p = point(15.0, 0.0, 7.0)


Scenario: The transformation matrix for the default orientation
  Given from ← point(0.0, 0.0, 0.0)
    And to ← point(0.0, 0.0, -1.0)
    And up ← vector(0.0, 1.0, 0.0)
  When t ← view_transform(from, to, up)
  Then t = identity_matrix

Scenario: A view transformation matrix looking in positive z direction
  Given from ← point(0.0, 0.0, 0.0)
    And to ← point(0.0, 0.0, 1.0)
    And up ← vector(0.0, 1.0, 0.0)
  When t ← view_transform(from, to, up)
  Then t = scaling(-1.0, 1.0, -1.0)

Scenario: The view transformation moves the world
  Given from ← point(0.0, 0.0, 8.0)
    And to ← point(0.0, 0.0, 0.0)
    And up ← vector(0.0, 1.0, 0.0)
  When t ← view_transform(from, to, up)
  Then t = translation(0.0, 0.0, -8.0)

Scenario: An arbitrary view transformation
  Given from ← point(1.0, 3.0, 2.0)
    And to ← point(4.0, -2.0, 8.0)
    And up ← vector(1.0, 1.0, 0.0)
  When t ← view_transform(from, to, up)
  Then t is the following 4x4 matrix:
      | -0.50709 | 0.50709 |  0.67612 | -2.36643 |
      |  0.76772 | 0.60609 |  0.12122 | -2.82843 |
      | -0.35857 | 0.59761 | -0.71714 |  0.00000 |
      |  0.00000 | 0.00000 |  0.00000 |  1.00000 |
