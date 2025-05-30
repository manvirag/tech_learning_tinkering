package main

import (
	"fmt"
)

func main() {
	// Create a quadtree covering the NYC area
	// NYC roughly spans from 40.5°N to 40.9°N and 73.7°W to 74.0°W
	qt := NewQuadTree(40.5, -74.0, 0.4, 0.3, 4)

	// Add some sample restaurants with realistic NYC coordinates
	restaurants := []Restaurant{
		{ID: "1", Name: "Pizza Place", Location: Point{X: 40.7128, Y: -73.9857}, Cuisine: "Italian"},   // Times Square
		{ID: "2", Name: "Sushi Bar", Location: Point{X: 40.7589, Y: -73.9851}, Cuisine: "Japanese"},    // Midtown
		{ID: "3", Name: "Burger Joint", Location: Point{X: 40.7282, Y: -73.7949}, Cuisine: "American"}, // Queens
		{ID: "4", Name: "Taco Stand", Location: Point{X: 40.6782, Y: -73.9442}, Cuisine: "Mexican"},    // Brooklyn
		{ID: "5", Name: "Curry House", Location: Point{X: 40.7580, Y: -73.9855}, Cuisine: "Indian"},    // Midtown
	}

	// Insert restaurants into the quadtree
	for _, restaurant := range restaurants {
		qt.Insert(restaurant)
		fmt.Printf("Added restaurant: %s at location (%.4f, %.4f)\n",
			restaurant.Name, restaurant.Location.X, restaurant.Location.Y)
	}

	// Find restaurants near Times Square (40.7580, -73.9855)
	searchPoint := Point{X: 40.7580, Y: -73.9855}
	radius := 0.01 // approximately 1.1 km radius

	fmt.Printf("\nSearching for restaurants near Times Square (%.4f, %.4f) within radius %.4f:\n",
		searchPoint.X, searchPoint.Y, radius)

	nearby := qt.FindNearby(searchPoint, radius)
	for _, restaurant := range nearby {
		fmt.Printf("- %s (%.4f, %.4f) - %s\n",
			restaurant.Name, restaurant.Location.X, restaurant.Location.Y, restaurant.Cuisine)
	}
}
