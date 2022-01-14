Feature: Groups

Scenario: Creating a new group
  Given g ← group()
  Then g.transform = identity_matrix
    And g is empty

Scenario: Adding a child to a group
  Given g ← group()
    And s ← test_shape()
  When add_child(g, s)
  Then g is not empty
    And g includes s
    And s.parent = g
