package main

/*
func (samples *SamplesList) Process(configuration *Configuration) (results []string) {
	for _, p := range *samples {
		if p.AB > configuration.MinAtBats {
			if p.WOBA >= configuration.TopwOBP && p.WAR >= configuration.GoodWAR {
				if p.Level == configuration.LeagueLevel {
					results = append(results, fmt.Sprintf(GREEN+"%s appears ready to advance. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
				}
				if configuration.OneLevelDown != "none" && p.Level == configuration.OneLevelDown {
					results = append(results, fmt.Sprintf(GREEN+"%s should be moved higher than %s immediately"+RESET, p.NameAndPosition(), configuration.LeagueLevel))
				}
				if configuration.OneLevelUp != "none" && p.Level == configuration.OneLevelUp {
					results = append(results, fmt.Sprintf(GREEN+"%s is where he should be in %s"+RESET, p.NameAndPosition(), configuration.OneLevelUp))
				}
			}
			if p.WOBA >= configuration.MedianwOBP && (p.WOBA < configuration.TopwOBP && p.WAR > 0) {
				if p.Age > configuration.MaxAge {
					results = append(results, fmt.Sprintf(YELLOW+"%s is too old but playing at or above level. Should be moved up. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
				}
				if p.Age <= configuration.MaxAge && p.Age >= configuration.CalcedGettingOld {
					results = append(results, fmt.Sprintf("%s is reaching max age but playing at the right level. %s", p.NameAndPosition(), p.StatsLine()))
				}
				if p.Age < configuration.CalcedGettingOld {
					results = append(results, fmt.Sprintf("%s is playing where he should be. %s", p.NameAndPosition(), p.StatsLine()))
				}
			}
			if p.WOBA < configuration.MedianwOBP && (p.WOBA >= configuration.BottomwOBP && p.WAR >= configuration.MinWAR) {
				if p.Age > configuration.MaxAge {
					results = append(results, fmt.Sprintf(RED+"%s is too old and not playing well at this level. He should be cut. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
				} else {
					if configuration.OneLevelDown != "none" && p.Level == configuration.OneLevelDown {
						results = append(results, fmt.Sprintf(BLUE+"%s has not played well at %s and is at the right level now. %s"+RESET, p.NameAndPosition(), configuration.LeagueLevel, p.StatsLine()))
					} else {
						if p.Age <= configuration.MaxAge && p.Age >= configuration.CalcedGettingOld {
							if p.Potential >= configuration.MinStars {
								results = append(results, fmt.Sprintf(YELLOW+"%s is reaching max age and not playing well at this level. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
							}
							if p.Potential < configuration.MinStars {
								results = append(results, fmt.Sprintf(RED+"%s is reaching max age and not playing well at this level. He should be cut. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
							}
						}
						if p.Age < configuration.CalcedGettingOld {
							results = append(results, fmt.Sprintf(YELLOW+"%s is young and not playing well. Possibly drop a level until skillset grows a bit. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
						}
					}
				}
			}
			if p.WOBA < configuration.BottomwOBP || p.WAR < configuration.MinWAR {
				if p.Age >= configuration.CalcedGettingOld && p.Age <= configuration.MaxAge {
					if p.Potential >= configuration.MinStars {
						results = append(results, fmt.Sprintf(BLUE+"%s is playing horribly and needs to be dropped a level. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
					}
					if p.Potential < configuration.MinStars {
						results = append(results, fmt.Sprintf(RED+"%s is reaching max age and playing horribly. He should be cut. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
					}
				}
				if p.Age < configuration.CalcedGettingOld {
					results = append(results, fmt.Sprintf(BLUE+"%s is young and playing horribly. Drop a level to grow skills. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
				}
			}
		}
	}

	return
}
*/
