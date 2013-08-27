package main

import (
//"fmt"
)

type LeagueSample struct {
	League                             string
	AB, HR, OPSPlus                    int64
	AVG, OBP, SLG, WOBA, VORP, WAR, ZR float64
}

type SamplesList []*LeagueSample
