package main

import (
	"math"
)

type Point struct {
	X, Y float64
}

type Restaurant struct {
	ID       string
	Name     string
	Location Point
	Cuisine  string
}

type QuadTree struct {
	Boundary struct {
		X, Y, Width, Height float64
	}
	Capacity  int
	Points    []Restaurant
	Divided   bool
	NorthWest *QuadTree
	NorthEast *QuadTree
	SouthWest *QuadTree
	SouthEast *QuadTree
}

func NewQuadTree(x, y, width, height float64, capacity int) *QuadTree {
	return &QuadTree{
		Boundary: struct {
			X, Y, Width, Height float64
		}{
			X: x, Y: y, Width: width, Height: height,
		},
		Capacity: capacity,
		Points:   make([]Restaurant, 0),
		Divided:  false,
	}
}

func (qt *QuadTree) Insert(restaurant Restaurant) bool {
	if !qt.contains(restaurant.Location) {
		return false
	}

	if len(qt.Points) < qt.Capacity && !qt.Divided {
		qt.Points = append(qt.Points, restaurant)
		return true
	}

	if !qt.Divided {
		qt.subdivide()
	}

	return qt.NorthWest.Insert(restaurant) ||
		qt.NorthEast.Insert(restaurant) ||
		qt.SouthWest.Insert(restaurant) ||
		qt.SouthEast.Insert(restaurant)
}

func (qt *QuadTree) Query(rangeX, rangeY, rangeWidth, rangeHeight float64) []Restaurant {
	var found []Restaurant

	if !qt.intersects(rangeX, rangeY, rangeWidth, rangeHeight) {
		return found
	}

	for _, point := range qt.Points {
		if point.Location.X >= rangeX && point.Location.X <= rangeX+rangeWidth &&
			point.Location.Y >= rangeY && point.Location.Y <= rangeY+rangeHeight {
			found = append(found, point)
		}
	}

	if qt.Divided {
		found = append(found, qt.NorthWest.Query(rangeX, rangeY, rangeWidth, rangeHeight)...)
		found = append(found, qt.NorthEast.Query(rangeX, rangeY, rangeWidth, rangeHeight)...)
		found = append(found, qt.SouthWest.Query(rangeX, rangeY, rangeWidth, rangeHeight)...)
		found = append(found, qt.SouthEast.Query(rangeX, rangeY, rangeWidth, rangeHeight)...)
	}

	return found
}

func (qt *QuadTree) FindNearby(point Point, radius float64) []Restaurant {
	rangeX := point.X - radius
	rangeY := point.Y - radius
	rangeWidth := radius * 2
	rangeHeight := radius * 2

	restaurants := qt.Query(rangeX, rangeY, rangeWidth, rangeHeight)
	var nearby []Restaurant

	for _, restaurant := range restaurants {
		distance := math.Sqrt(
			math.Pow(restaurant.Location.X-point.X, 2) +
				math.Pow(restaurant.Location.Y-point.Y, 2))
		if distance <= radius {
			nearby = append(nearby, restaurant)
		}
	}

	return nearby
}

func (qt *QuadTree) contains(point Point) bool {
	return point.X >= qt.Boundary.X &&
		point.X <= qt.Boundary.X+qt.Boundary.Width &&
		point.Y >= qt.Boundary.Y &&
		point.Y <= qt.Boundary.Y+qt.Boundary.Height
}

func (qt *QuadTree) intersects(x, y, width, height float64) bool {
	return !(x+width < qt.Boundary.X ||
		x > qt.Boundary.X+qt.Boundary.Width ||
		y+height < qt.Boundary.Y ||
		y > qt.Boundary.Y+qt.Boundary.Height)
}

func (qt *QuadTree) subdivide() {
	x := qt.Boundary.X
	y := qt.Boundary.Y
	w := qt.Boundary.Width / 2
	h := qt.Boundary.Height / 2

	qt.NorthWest = NewQuadTree(x, y, w, h, qt.Capacity)
	qt.NorthEast = NewQuadTree(x+w, y, w, h, qt.Capacity)
	qt.SouthWest = NewQuadTree(x, y+h, w, h, qt.Capacity)
	qt.SouthEast = NewQuadTree(x+w, y+h, w, h, qt.Capacity)

	qt.Divided = true
}
