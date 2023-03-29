package main

import "fmt"

type (
	// Scenario It is the environment where the player interacts. It contains the rules and restrictions that dictate
	// the gamification dynamics.
	Scenario struct {
		ID              ScenarioID
		GameMode        GameMode
		Site            Site Game            Game
		VehicleCategory CategoryID
	}

	Scenarios  []Scenario
	Site       string
	GameMode   string
	Game       string
	CategoryID string
	ScenarioID string

	Filter interface {
		IsSatisfied(s Scenario) bool
	}
	Filters    []Filter
	SiteFilter Site
	ModeFilter GameMode
	IDFilter   ScenarioID
)

func (f Filters) IsSatisfied(s Scenario) bool {
	for _, ff := range f {
		if !ff.IsSatisfied(s) {
			return false
		}
	}

	return true
}

func (s SiteFilter) IsSatisfied(sc Scenario) bool {
	return Site(s) == sc.Site
}

func (m ModeFilter) IsSatisfied(sc Scenario) bool {
	return GameMode(m) == sc.GameMode
}

func (i IDFilter) IsSatisfied(sc Scenario) bool {
	return ScenarioID(i) == sc.ID
}

func (s Scenarios) Filter(f Filters) Scenarios {
	r := make(Scenarios, 0, len(s))

	for _, sc := range s {
		if f.IsSatisfied(sc) {
			r = append(r, sc)
		}
	}

	return r
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
	sSite := SiteFilter("MLZ")
	sID := IDFilter("1")

	r := scenarios.Filter(Filters{sSite, sID})

	fmt.Println(r)
}
