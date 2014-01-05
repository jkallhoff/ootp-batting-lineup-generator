package main

import (
//"fmt"
)

type LeagueSample struct {
	League                                                                               string
	LeagueInt                                                                            int64
	AB, HR, OPSPlus                                                                      int64
	AVG, OBP, SLG, WOBA, VORP, WAR, ZR                                                   float64
	AtAge, GettingOld, OverAge                                                           bool
	PlayingHorribly, PlayingBelowAverage, PlayingAverage, PlayingAboveAverage, KillingIt bool
	PoorDefense                                                                          bool
	Processed                                                                            bool
}

type SamplesList []*LeagueSample
