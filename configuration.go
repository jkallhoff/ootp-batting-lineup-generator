package main

import (
	"github.com/JKallhoff/gofig"
)

type Configuration struct {
	MinAtBats, AMaxAge, AAMaxAge, AAAMaxAge, MLMaxAge, MinStars, CalcedGettingOld int64
	MedianwOBP, MinWAR, GoodWAR, BottomwOBP, TopwOBP                              float64
}

func (configuration *Configuration) Load() {
	if conf, err := gofig.Load("./gofig.json"); err == nil {
		configuration.MinAtBats, _ = conf.Int64("minAtBats")
		configuration.AMaxAge, _ = conf.Int64("amaxAge")
		configuration.AAMaxAge, _ = conf.Int64("aamaxAge")
		configuration.AAAMaxAge, _ = conf.Int64("aaamaxAge")
		configuration.MinStars, _ = conf.Int64("minStars")
		configuration.MedianwOBP, _ = conf.Float("medianwOBP")
		configuration.MinWAR, _ = conf.Float("minWAR")
		configuration.GoodWAR, _ = conf.Float("goodWAR")
		configuration.BottomwOBP, _ = conf.Float("bottomwOBP")
		configuration.TopwOBP, _ = conf.Float("topwOBP")
	}
}
