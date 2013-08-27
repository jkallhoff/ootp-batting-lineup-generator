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

/*
func (sample *LeagueSample) NameAndPosition() string {
	return fmt.Sprintf("%s(%s %s %v stars)", sample.Name, sample.Position, sample.Level, sample.Potential)
}

func (sample *LeagueSample) StatsLine() string {
	return fmt.Sprintf("wOBA: %0.3f WAR: %0.2f AB: %v", sample.WOBA, sample.WAR, sample.AB)
}
*/
