Feature: Materials

Background:
  Given m ← material()
    And position ← point(0.0, 0.0, 0.0)

Scenario: The default material
  Given m ← material()
  Then m.color = color(1.0, 1.0, 1.0)
    And m.ambient = 0.1
    And m.diffuse = 0.9
    And m.specular = 0.9
    And m.shininess = 200.0

Scenario: Lighting with the eye between the light and the surface
  Given eyev ← vector(0.0, 0.0, -1.0)
    And normalv ← vector(0.0, 0.0, -1.0)
    And light ← point_light(point(0.0, 0.0, -10.0), color(1.0, 1.0, 1.0))
  When result ← lighting(m, light, position, eyev, normalv)
  Then result = color(1.9, 1.9, 1.9)
