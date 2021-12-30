Feature: Lights

Scenario: A point light has a position and intensity
  Given intensity ← color(1.0, 1.0, 1.0)
    And position ← point(0.0, 0.0, 0.0)
  When light ← point_light(position, intensity)
  Then light.position = position
    And light.intensity = intensity
