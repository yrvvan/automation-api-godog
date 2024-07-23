package main

import (
	"os"
	"testing"

	"godog-automation/steps"

	"github.com/cucumber/godog"
)

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^there are no items$`, steps.ThereAreNoItems)
	s.Step(`^I create (\d+) item named "([^"]*)"$`, steps.ICreateAnItemNamed)
	s.Step(`^I should have (\d+) items$`, steps.IShouldHaveItems)
	s.Step(`^an item named "([^"]*)"$`, steps.AnItemNamed)
	s.Step(`^I read the item with ID (\d+)$`, steps.IReadTheItemWithID)
	s.Step(`^I should get the item with name "([^"]*)"$`, steps.IShouldGetTheItemWithName)
	s.Step(`^I update the item with ID (\d+) to have name "([^"]*)"$`, steps.IUpdateTheItemWithIDToHaveName)
	s.Step(`^the item with ID (\d+) should have name "([^"]*)"$`, steps.TheItemWithIDShouldHaveName)
	s.Step(`^I delete the item with ID (\d+)$`, steps.IDeleteTheItemWithID)
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "cucumber",
		Paths:  []string{"features"},
		Output: os.Stdout,
	}
	status := godog.TestSuite{
		Name:                "godog",
		ScenarioInitializer: FeatureContext,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
