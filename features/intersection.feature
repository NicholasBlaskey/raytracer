Feature: Intersections

Scenario: An intersection encapsulates t and object
  Given s ← sphere()
  When i ← intersection(3.5, s)
  Then i.t = 3.5
    And i.object = s

Scenario: Aggregating intersections
  Given s ← sphere()
    And i1 ← intersection(1.0, s)
    And i2 ← intersection(2.0, s)
  When xs ← intersections(i1, i2)
  Then xs.count = 2
    And xs[0].t = 1.0
    And xs[1].t = 2.0


Scenario: The hit, when all intersections have positive t
  Given s ← sphere()
    And i1 ← intersection(1.0, s)
    And i2 ← intersection(2.0, s)
    And xs ← intersections(i2, i1)
  When i ← hit(xs)
  Then i = i1

Scenario: The hit, when some intersections have negative t
  Given s ← sphere()
    And i1 ← intersection(-1.0, s)
    And i2 ← intersection(1.0, s)
    And xs ← intersections(i2, i1)
  When i ← hit(xs)
  Then i = i2

Scenario: The hit, when all intersections have negative t
  Given s ← sphere()
    And i1 ← intersection(-2.0, s)
    And i2 ← intersection(-1.0, s)
    And xs ← intersections(i2, i1)
  When i ← hit(xs)
  Then i is nothing

Scenario: The hit is always the lowest nonnegative intersection
  Given s ← sphere()
  And i1 ← intersection(5.0, s)
  And i2 ← intersection(7.0, s)
  And i3 ← intersection(-3.0, s)
  And i4 ← intersection(2.0, s)
  And xs ← intersections(i1, i2, i3, i4)
When i ← hit(xs)
Then i = i4

Scenario: Precomputing the state of an intersection
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And shape ← sphere()
    And i ← intersection(4.0, shape)
  When comps ← prepare_computations(i, r)
  Then comps.t = i.t
    And comps.object = i.object
    And comps.point = point(0.0, 0.0, -1.0)
    And comps.eyev = vector(0.0, 0.0, -1.0)
    And comps.normalv = vector(0.0, 0.0, -1.0)

Scenario: Precomputing the reflection vector
  Given shape ← plane()
    And r ← ray(point(0.0, 1.0, -1.0), vector(0.0, -0.70710678118, 0.70710678118)) 
    And i ← intersection(1.41421356237, shape)                      
  When comps ← prepare_computations(i, r)
  Then comps.reflectv = vector(0.0, 0.70710678118, 0.70710678118)                

Scenario: The hit, when an intersection occurs on the outside
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And shape ← sphere()
    And i ← intersection(4.0, shape)
  When comps ← prepare_computations(i, r)
  Then comps.inside = false

Scenario: The hit, when an intersection occurs on the inside
  Given r ← ray(point(0.0, 0.0, 0.0), vector(0.0, 0.0, 1.0))
    And shape ← sphere()
    And i ← intersection(1.0, shape)
  When comps ← prepare_computations(i, r)
  Then comps.point = point(0.0, 0.0, 1.0)
    And comps.eyev = vector(0.0, 0.0, -1.0)
    And comps.inside = true
      # normal would have been (0, 0, 1), but is inverted!
    And comps.normalv = vector(0.0, 0.0, -1.0)

Scenario: The hit should offset the point
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And shape ← sphere() with:
      | transform | translation(0, 0, 1) |
    And i ← intersection(5.0, shape)
  When comps ← prepare_computations(i, r)
  Then comps.over_point.z < -EPSILON/2
    And comps.point.z > comps.over_point.z

Scenario Outline: Finding n1 and n2 at various intersections
  Given A ← glass_sphere() with:
      | transform                 | scaling(2, 2, 2) |
      | material.refractive_index | 1.5              |
    And B ← glass_sphere() with:
      | transform                 | translation(0, 0, -0.25) |
      | material.refractive_index | 2.0                      |
    And C ← glass_sphere() with:
      | transform                 | translation(0, 0, 0.25) |
      | material.refractive_index | 2.5                     |
    And r ← ray(point(0.0, 0.0, -4.0), vector(0.0, 0.0, 1.0))
    And xs ← intersections(2.0:A, 2.75:B, 3.25:C, 4.75:B, 5.25:C, 6:A)
  When comps ← prepare_computations(xs[<index>], r, xs)  
  Then comps.n1 = <n1>
    And comps.n2 = <n2>             

  Examples:
    | index | n1  | n2  |
    | 0     | 1.0 | 1.5 |                 
    | 1     | 1.5 | 2.0 |
    | 2     | 2.0 | 2.5 |
    | 3     | 2.5 | 2.5 |
    | 4     | 2.5 | 1.5 |
    | 5     | 1.5 | 1.0 |

Scenario: The under point is offset below the surface
  Given r ← ray(point(0.0, 0.0, -5.0), vector(0.0, 0.0, 1.0))
    And shape ← glass_sphere() with:
      | transform | translation(0, 0, 1) |
    And i ← intersection(5.0, shape)
    And xs ← intersections(i)
  When comps ← prepare_computations(i, r, xs)
  Then comps.under_point.z > EPSILON/2
    And comps.point.z < comps.under_point.z
