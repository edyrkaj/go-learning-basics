# features/user_login.feature
Feature: User Login
  Scenario: Successful login
    Given a registered user "alice" with password "password123"
    When I submit the login request with "alice" and "password123"
    Then the response status code should be 200
