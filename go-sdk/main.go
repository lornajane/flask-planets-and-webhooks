package main

import (
	"context"
	"fmt"

	"github.com/lornajane/flask-planets-and-webhooks/go-sdk/planets"
)

func AllPlanets() []planets.Planet {
	client := planets.NewAPIClient(planets.NewConfiguration())
	ctx := context.Background()
	planets, _, err := client.PlanetsApi.AllPlanets(ctx)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	return planets
}

func OnePlanet(planetId float32) planets.Planet {
	client := planets.NewAPIClient(planets.NewConfiguration())
	ctx := context.Background()
	planet, _, err := client.PlanetsApi.OnePlanet(ctx, planetId)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	return planet
}
