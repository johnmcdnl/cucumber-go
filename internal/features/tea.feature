Feature: Hello World

  Scenario: Tea 1
    Given some tea
    When I drink some tea
    Then I have no tea

  Scenario: Tea 2
    Given lots of tea
    When I drink some
    Then I have some tea left

  Scenario: Tea 3
    Given some tea
    When I try to drink lots of tea
    Then I don't have enough