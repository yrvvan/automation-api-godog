Feature: CRUD operations

    Scenario: Create an item
        Given there are no items
        When I create 2 item named "example"
        Then I should have 2 items
        And an item named "example"

    Scenario: Read an item
        Given I create 1 item named "example"
        When I read the item with ID 1
        Then I should get the item with name "example"

    Scenario: Update an item
        Given I create 10 item named "example"
        When I update the item with ID 10 to have name "new_example"
        Then the item with ID 10 should have name "new_example"

    Scenario: Delete an item
        Given I create 2 item named "example"
        When I delete the item with ID 1
        Then I should have 1 items
