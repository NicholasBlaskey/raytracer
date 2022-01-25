Feature: Constructive Solid Geometry (CSG)

Scenario: CSG is created with an operation and two shapes
  Given s1 ← sphere()
    And s2 ← cube()
  When c ← csg("union", s1, s2)
  Then c.operation = "union"
    And c.left = s1
    And c.right = s2
    And s1.parent = c
    And s2.parent = c

Scenario Outline: Evaluating the rule for a CSG operation
  When result ← intersection_allowed("<op>", <lhit>, <inl>, <inr>)
  Then result = <result>

  Examples:
  | op           | lhit  | inl   | inr   | result |
  | union        | true  | true  | true  | false  |
  | union        | true  | true  | false | true   |
  | union        | true  | false | true  | false  |
  | union        | true  | false | false | true   |
  | union        | false | true  | true  | false  |
  | union        | false | true  | false | false  |
  | union        | false | false | true  | true   |
  | union        | false | false | false | true   |
  #
  | intersection | true  | true  | true  | true   |
  | intersection | true  | true  | false | false  |
  | intersection | true  | false | true  | true   |
  | intersection | true  | false | false | false  |
  | intersection | false | true  | true  | true   |
  | intersection | false | true  | false | true   |
  | intersection | false | false | true  | false  |
  | intersection | false | false | false | false  |
  #
  | difference   | true  | true  | true  | false  |
  | difference   | true  | true  | false | true   |
  | difference   | true  | false | true  | false  |
  | difference   | true  | false | false | true   |
  | difference   | false | true  | true  | true   |
  | difference   | false | true  | false | true   |
  | difference   | false | false | true  | false  |
  | difference   | false | false | false | false  |
