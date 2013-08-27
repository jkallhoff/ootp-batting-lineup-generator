package main

type Player struct {
	Position, Name, Level, Bats                                         string
	Age, Contact, Gap, Power, Eye, AvoidsStrikeOuts, Overall, Potential int64
	Catcher, FirstBase, SecondBase, ThirdBase                           int64
	ShortStop, LeftField, RightField, CenterField                       int64
	Speed, Stealing, BaseRunning                                        int64
	Samples                                                             SamplesList
}
