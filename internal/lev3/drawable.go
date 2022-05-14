package lev3

/*Integrace that could be drawn by printer.*/
type Drawable interface {

	GetStartPoint() Point

	GetStepsCount() int

	GetStep(index int) Point

	GetStepDiff(index int) Point
}
