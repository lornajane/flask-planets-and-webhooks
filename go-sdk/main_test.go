package main

import (
	"testing"

	"github.com/lornajane/flask-planets-and-webhooks/go-sdk/planets"
)

func TestAllPlanets(t *testing.T) {
	all_the_planets := AllPlanets()

	good_example := planets.Planet{Name: "Uranus", Position: 7, Moons: 27}
	if all_the_planets[6] != good_example {
		t.Errorf("All planets returned an unexpected response")
	}
}

func TestOnePlanet(t *testing.T) {
	one_planet := OnePlanet(2)

	good_example := planets.Planet{Name: "Venus", Position: 2, Moons: 0}
	if one_planet != good_example {
		t.Errorf("One planet went wrong")
	}

}
