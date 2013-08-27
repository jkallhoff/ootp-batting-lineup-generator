package main

import (
	"github.com/jkallhoff/gofig"
)

type Configuration struct {
	LeagueLevel, OneLevelDown, OneLevelUp            string
	MinAtBats, MaxAge, MinStars, CalcedGettingOld    int64
	MedianwOBP, MinWAR, GoodWAR, BottomwOBP, TopwOBP float64
}

func (configuration *Configuration) Load() {
	configuration.LeagueLevel = gofig.Str("leagueLevel")
	configuration.MinAtBats = gofig.Int("minAtBats")
	configuration.MaxAge = gofig.Int("maxAge")
	configuration.MinStars = gofig.Int("minStars")
	configuration.MedianwOBP = gofig.Float("medianwOBP")
	configuration.MinWAR = gofig.Float("minWAR")
	configuration.GoodWAR = gofig.Float("goodWAR")
	configuration.BottomwOBP = gofig.Float("bottomwOBP")
	configuration.TopwOBP = gofig.Float("topwOBP")

	configuration.CalcedGettingOld = configuration.MaxAge - 2

	switch configuration.LeagueLevel {
	case "A":
		configuration.OneLevelDown = "none"
		configuration.OneLevelUp = "AA"
	case "AA":
		configuration.OneLevelDown = "A"
		configuration.OneLevelUp = "AAA"
	case "AAA":
		configuration.OneLevelDown = "AA"
		configuration.OneLevelUp = "ML"
	case "ML":
		configuration.OneLevelDown = "AAA"
		configuration.OneLevelUp = "none"
	}
}
