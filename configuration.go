package main

import (
	"github.com/jkallhoff/gofig"
)

type Configuration struct {
	MinAtBats, AMaxAge, AAMaxAge, AAAMaxAge, MLMaxAge, MinStars, CalcedGettingOld int64
	MedianwOBP, MinWAR, GoodWAR, BottomwOBP, TopwOBP                              float64
}

func (configuration *Configuration) Load() {
	configuration.MinAtBats = gofig.Int("minAtBats")
	configuration.AMaxAge = gofig.Int("amaxAge")
	configuration.AAMaxAge = gofig.Int("aamaxAge")
	configuration.AAAMaxAge = gofig.Int("aaamaxAge")
	configuration.MinStars = gofig.Int("minStars")
	configuration.MedianwOBP = gofig.Float("medianwOBP")
	configuration.MinWAR = gofig.Float("minWAR")
	configuration.GoodWAR = gofig.Float("goodWAR")
	configuration.BottomwOBP = gofig.Float("bottomwOBP")
	configuration.TopwOBP = gofig.Float("topwOBP")
}
