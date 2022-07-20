package greetings

import "time"

func GoodDay(name string) []string {
	s := []string{"GoodDay"}
	g := append(s, name)
	return g
}
func GoodNight(name string) []string {
	s := []string{"GoodNight"}
	g := append(s, name)
	return g
}

var now = time.Now()
var nowhour = now.Hour()

func IsAM() bool {
	if nowhour >= 5 && nowhour < 12 {
		return true
	} else {
		return false
	}
}
func IsAfternoon() bool {
	if nowhour >= 12 && nowhour < 18 {
		return true
	} else {
		return false
	}
}
func IsEvening() bool {
	if nowhour >= 18 && nowhour < 21 {
		return true
	} else {
		return false
	}
}
