Feature: Abstract Shapes

Scenario: The default transformation
  Given s ← test_shape()
  Then s.transform = identity_matrix

Scenario: Assigning a transformation
  Given s ← test_shape()
  When set_transform(s, translation(2.0, 3.0, 4.0))
  Then s.transform = translation(2.0, 3.0, 4.0)
