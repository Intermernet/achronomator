package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

var (
	decadeStart = 1200
	decadeEnd   = 2050

	descriptions = []string{
		"Sparkly",
		"Stealthy",
		"Mighty",
		"Polkadotted",
		"Plaid",
		"Itinerant",
		"Beaded",
		"Bejewelled",
		"Future",
		"Mystical",
		"Lurid",
		"Monochromatic",
		"Eccentric",
		"Vibrant",
		"Disco",
		"Shaggy",
		"Floral",
		"Refined",
		"Fraggly",
		"Avant Garde",
		"Sophisticated",
		"Philanthropic",
		"Courageous",
		"Feathery",
		"Fearless",
		"Robotic",
		"Metallic",
		"Hypercolour",
		"Cheesy",
	}

	jobs = []string{
		"Pirate",
		"Cosmonaut",
		"Salsa Dancer",
		"Zeppelin Pilot",
		"Surgeon",
		"Explorer",
		"Flapper",
		"Detective",
		"Snake Wrangler",
		"Sailor",
		"Ninja",
		"Plumber",
		"Pool Cleaner",
		"Musician",
		"Smuggler",
		"Admiral",
		"Fencer",
		"Dancer",
		"Superhero",
		"Scuba Diver",
		"Rockstar",
		"Scientist",
		"Bounty Hunter",
		"Alchemist",
		"Hermit",
		"Fraggle",
		"Zoo Keeper",
		"Lion Tamer",
		"Batman",
		"Trapeeze Artist",
		"Raver",
	}
)

func main() {
	//fmt.Printf(acronomate())
	js.Global.Get("document").Call("write", acronomate())
}

func acronomate() string {
	rand.Seed(time.Now().UnixNano())
	decades := make([]int, 0)
	for decade := decadeStart; decade <= decadeEnd; decade += 10 {
		decades = append(decades, decade)
	}
	era := decades[rand.Intn(len(decades))]
	description := descriptions[rand.Intn(len(descriptions))]
	job := jobs[rand.Intn(len(jobs))]
	return fmt.Sprintf("%d's %s %s\n", era, description, job)
}
