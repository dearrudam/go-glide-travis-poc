Feature: run calculation features

  Scenario Outline: should run a addition calculation
    When I sum <num1> with <num2>
    Then the result should be <result>

    Examples:
      | num1 | num2 | result |
      | 1    | 1    | 2      |