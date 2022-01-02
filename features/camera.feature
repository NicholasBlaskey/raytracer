Feature: Camera

Scenario: Constructing a camera
  Given hsize ← 160
    And vsize ← 120
    And field_of_view ← 1.57079632679
  When c ← camera(hsize, vsize, field_of_view)
  Then c.hsize = 160
    And c.vsize = 120
    And c.field_of_view = 1.57079632679
    And c.transform = identity_matrix
