package types

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type ListPoint struct {
	Points []Point
}

// 1,1:2,0:2,3:3,4:4,3
// turn into list of Points
func (l *ListPoint) PopulateListPoint(input string) {
	parsedInput := strings.Split(input, ":")
	newPointsArray := []Point{}
	for _, c := range parsedInput {
		newPoint := Point{}
		loc := strings.Split(string(c), ",")
		x, _ := strconv.Atoi(loc[0])
		y, _ := strconv.Atoi(loc[1])
		newPoint.X = x
		newPoint.Y = y
		newPointsArray = append(newPointsArray, newPoint)
	}
	l.Points = newPointsArray
}
