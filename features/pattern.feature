Feature: Patterns

Background:
  Given black ← color(0.0, 0.0, 0.0)
    And white ← color(1.0, 1.0, 1.0)

Scenario: Creating a stripe pattern
  Given pattern ← stripe_pattern(white, black)
  Then pattern.a = white
    And pattern.b = black

Scenario: A stripe pattern is constant in y
  Given pattern ← stripe_pattern(white, black)
  Then stripe_at(pattern, point(0.0, 0.0, 0.0)) = white
    And stripe_at(pattern, point(0.0, 1.0, 0.0)) = white
    And stripe_at(pattern, point(0.0, 2.0, 0.0)) = white

Scenario: A stripe pattern is constant in z
  Given pattern ← stripe_pattern(white, black)
  Then stripe_at(pattern, point(0.0, 0.0, 0.0)) = white
    And stripe_at(pattern, point(0.0, 0.0, 1.0)) = white
    And stripe_at(pattern, point(0.0, 0.0, 2.0)) = white

Scenario: A stripe pattern alternates in x
  Given pattern ← stripe_pattern(white, black)
  Then stripe_at(pattern, point(0.0, 0.0, 0.0)) = white
    And stripe_at(pattern, point(0.9, 0.0, 0.0)) = white
    And stripe_at(pattern, point(1.0, 0.0, 0.0)) = black
    And stripe_at(pattern, point(-0.1, 0.0, 0.0)) = black
    And stripe_at(pattern, point(-1.0, 0.0, 0.0)) = black
    And stripe_at(pattern, point(-1.1, 0.0, 0.0)) = white
