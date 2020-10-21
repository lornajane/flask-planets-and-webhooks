/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/lornajane/flask-planets-and-webhooks/goplanet"
	"github.com/spf13/cobra"
)

var Position int

// planetCmd represents the planet command
var planetCmd = &cobra.Command{
	Use:   "planet",
	Short: "Work with planets",
	Long:  `Fetch or list planets`,
	Run: func(cmd *cobra.Command, args []string) {
		if Position > 0 {
			planet := goplanet.OnePlanet(float32(Position))
			moons := fmt.Sprintf("%g", planet.Moons)
			fmt.Println(planet.Name + " with " + moons + " moons")
		} else {
			list := goplanet.AllPlanets()

			for _, planet := range list {
				moons := fmt.Sprintf("%g", planet.Moons)
				fmt.Println(planet.Name + " with " + moons + " moons")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(planetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// planetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	planetCmd.Flags().IntVarP(&Position, "position", "p", 0, "Position of planet to select")
}
