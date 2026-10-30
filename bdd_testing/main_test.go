// main_test.go
package main

import (
	"context"

	"github.com/cucumber/godog"
)

func aRegisteredUserWithPassword(ctx context.Context, username, password string) error {
	// Setup user state in database
	return nil
}

func iSubmitTheLoginRequest(ctx context.Context, username, password string) error {
	// Execute HTTP call
	return nil
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^a registered user "([^"]*)" with password "([^"]*)"$`, aRegisteredUserWithPassword)
	sc.Step(`^I submit the login request with "([^"]*)" and "([^"]*)"$`, iSubmitTheLoginRequest)
}
