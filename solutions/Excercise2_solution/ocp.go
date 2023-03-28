package main

import "fmt"

type (
	// Scenario It is the environment where the player interacts. It contains the rules and restrictions that dictate
	// the gamification dynamics.
	Scenario struct {
		ID              ScenarioID
		GameMode        GameMode
		Site            Site
		Game            Game
		VehicleCategory CategoryID
	}

	Scenarios  []Scenario
	Site       string
	GameMode   string
	Game       string
	CategoryID string
	ScenarioID string

	SiteSpecification struct {
		Site Site
	}

	ModeSpecification struct {
		Mode GameMode
	}

	GameSpecification struct {
		Game Game
	}

	CategoryAndScenarioSpecification struct {
		Category CategoryID
		Scenario ScenarioID
	}

	Specification interface {
		IsSatisfied(p *Scenario) bool
	}
)

func (c CategoryAndScenarioSpecification) IsSatisfied(p *Scenario) bool {
	return p.VehicleCategory == c.Category && p.ID == c.Scenario
}

func (s SiteSpecification) IsSatisfied(p *Scenario) bool {
	return p.Site == s.Site
}

func (m ModeSpecification) IsSatisfied(p *Scenario) bool {
	return p.GameMode == m.Mode
}

func (g GameSpecification) IsSatisfied(p *Scenario) bool {
	return p.Game == g.Game
}

func (s Scenarios) Filter(specification Specification) Scenarios {
	result := make(Scenarios, 0)
	for i, v := range s {
		if specification.IsSatisfied(&v) {
			result = append(result, s[i])
		}
	}

	return result
}

func main() {
	scenarios := Scenarios{
		Scenario{
			ID:              "1",
			GameMode:        "crowd",
			Site:            "MLT",
			Game:            "1",
			VehicleCategory: "12",
		},
		Scenario{
			ID:              "2",
			GameMode:        "logistics",
			Site:            "MLF",
			Game:            "1",
			VehicleCategory: "11",
		},
		Scenario{
			ID:              "3",
			GameMode:        "crowd",
			Site:            "MLT",
			Game:            "2",
			VehicleCategory: "11",
		},
		Scenario{
			ID:              "1",
			GameMode:        "crowd",
			Site:            "MLZ",
			Game:            "2",
			VehicleCategory: "11",
		},
	}

	siteSpec := SiteSpecification{Site: "MLT"}
	modeSpec := ModeSpecification{Mode: "crowd"}
	gameSpec := GameSpecification{Game: "2"}
	catAndSceSpec := CategoryAndScenarioSpecification{Category: "11", Scenario: "2"}

	fmt.Println(scenarios.Filter(siteSpec))
	fmt.Println(scenarios.Filter(modeSpec))
	fmt.Println(scenarios.Filter(gameSpec))
	fmt.Println(scenarios.Filter(catAndSceSpec))
}
