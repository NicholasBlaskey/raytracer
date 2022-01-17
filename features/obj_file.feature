Feature: OBJ File Parser

Scenario: Ignoring unrecognized lines
  Given gibberish ← a file containing:
    """
    There was a young lady named Bright
    who traveled much faster than light.
    She set out one day
    in a relative way,
    and came back the previous night.
    """
  When parser ← parse_obj_file(gibberish)
  Then parser should have ignored 5 lines
