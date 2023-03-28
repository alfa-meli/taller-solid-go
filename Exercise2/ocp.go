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
)

// FilterBySite returns only the scenarios that match the site criteria.
func (s Scenarios) FilterBySite(site Site) Scenarios {
	scenarios := make(Scenarios, 0)

	for i := range s {
		if s[i].Site == site {
			scenarios = append(scenarios, s[i])
		}
	}

	return scenarios
}

// FilterByMode returns only the scenarios that match the game mode.
func (s Scenarios) FilterByMode(mode GameMode) Scenarios {
	scenarios := make(Scenarios, 0)

	for i := range s {
		if s[i].GameMode == mode {
			scenarios = append(scenarios, s[i])
		}
	}

	return scenarios
}

// FilterByGame returns only the scenarios that match the game.
func (s Scenarios) FilterByGame(game Game) Scenarios {
	scenarios := make(Scenarios, 0)

	for i := range s {
		if s[i].Game == game {
			scenarios = append(scenarios, s[i])
		}
	}

	return scenarios
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

	bySite := scenarios.FilterBySite("MLT")
	byMode := scenarios.FilterByMode("crowd")
	byGame := scenarios.FilterByGame("1")
	fmt.Println(bySite)
	fmt.Println(byMode)
	fmt.Println(byGame)
}
