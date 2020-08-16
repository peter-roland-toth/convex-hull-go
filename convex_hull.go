package main

import (
	"sort"
	// "math"
	// "fmt"
	// "encoding/json"
	// "io/ioutil"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Polygon struct {
	left_index int
	right_index int
	points []Point
}

// Returns the convex hull of the set of points provided in the input
func Calculate_hull(p []Point) []Point {
	points := preprocess(p)
	hull := convex_hull(points).points
	
	return hull
}

// Used to sort the points based on their X value and to remove the
// duplicates.
func preprocess(p []Point) []Point {
	// sorting
	sort.Slice(p, func(i, j int) bool {
		// if the X value is the same for two points, we order them
		// based on the Y value, to make sure that duplicate points
		// end up next to each other after sorting
		if p[i].X == p[j].X {
			return p[i].Y < p[j].Y
		}

		return p[i].X < p[j].X
	})

	result := []Point{p[0]}

	// adding only unique points to the result
	for i := 1; i < len(p); i += 1 {
		if p[i] != p[i-1] {
			result = append(result, p[i])
		}
	}

	return result
}

// Returns the convex hull for the points sent as parameter.
// The convex hull is the smallest polygon that cointains all
// the points sent to the method.
// The method assumes that the input is sorted along the X-axis
func convex_hull(p []Point) Polygon {
	if len(p) <= 1 {
		return Polygon{0, 0, p}
	}

	// partitioning the points into a left and right region
	left, right := partition(p)

	// creating the two convex hulls for the two regions
	poly_left := convex_hull(left)
	poly_right := convex_hull(right)

	// merging the two convex hulls into one big convex hull
	merged_polygon := merge(poly_left, poly_right)

	return merged_polygon
}

// Partitions the received points into two equal parts
func partition(p []Point) ([]Point, []Point) {
	median := len(p) / 2
	left := p[0:median]
	right := p[median:]

	return left, right
}

// Merges two convex hulls into one, using the "two-fingers algorithm"
// It assumes that the two input polygons are sorted clockwise
// The returned polygon is also sorted clockwise
// This algorithm has O(n) complexity
func merge(p1, p2 Polygon) Polygon {
	// if one of the input arrays is empty, we just return the other one
	if len(p1.points) == 0 {
		return p2
	}

	if len(p2.points) == 0 {
		return p1
	}

	// finding the rightmost point from the left region, 
	// and the leftmost point from the right region
	// max_i := -math.MaxFloat64
	// min_j := math.MaxFloat64
	// best_i, best_j := 0, 0

	// for i := 0; i < len(p1); i++ {
	// 	if p1[i].X > max_i {
	// 		max_i = p1[i].X
	// 		best_i = i
	// 	}
	// }

	// for j := 0; j < len(p2); j++ {
	// 	if p2[j].X < min_j {
	// 		min_j = p2[j].X
	// 		best_j = j
	// 	}
	// }

	i := p1.right_index
	j := p2.left_index
	max_i := p1.points[i].X
	min_j := p2.points[j].X
	

	// calculating the position of the vertical line between the two regions
	// it doesn't have to be exactly in the middle, but it's important that none
	// of the points are on this line. That could introduce problems when calculating
	// the upper and lower tangents below.
	x := (min_j + max_i) / 2.0

	// i, j := best_i, best_j

	// calculating the upper tangent between the two regions. The idea is that we start
	// from the closest points to the middle line calculated before, and going "upwards"
	// (going clockwise on the right side and counter-clockwise on the left side) we find
	// the point with the highest Y on the middle line.
	//
	//				 r2 X----
	//	    		   /
	//	---X l1       /
	//		\		 X  r1
 	//		 X l0    |
	//		/		 X  r0
	//  ---X ln       \
	//
	// in this case the algorithm would start from l0 and r0 respectively and would find
	// the l1-r2 line
	p1_points := p1.points
	p2_points := p2.points

	for {
		prev_i := (((i-1) % len(p1_points)) + len(p1_points)) % len(p1_points)
		next_j := (j+1) % len(p2_points)

		if y_value(p1_points[prev_i], p2_points[j], x) > y_value(p1_points[i], p2_points[j], x) {
			i = prev_i
		} else if y_value(p1_points[i], p2_points[next_j], x) > y_value(p1_points[i], p2_points[j], x) {
			j = next_j
		} else {
			break
		}
	}

	upper_i, upper_j := i, j
	i, j = p1.right_index, p2.left_index
	
	// same as above but here we find the lower tangent, going clockwise on the left side
	// now and counter-clockwise on the right side, to find the minimum Y value on the middle line
	for {
		next_i := (i+1) % len(p1_points)
		prev_j := (((j-1) % len(p2_points)) + len(p2_points)) % len(p2_points)

		if y_value(p1_points[next_i], p2_points[j], x) < y_value(p1_points[i], p2_points[j], x) {
			i = next_i
		} else if y_value(p1_points[i], p2_points[prev_j], x) < y_value(p1_points[i], p2_points[j], x) {
			j = prev_j
		} else {
			break
		}
	}

	lower_i, lower_j := i, j

	// merging the two regions. First we add the upper tangent, then we go clockwise on the right
	// side until we find the point where we can add the lower tangent. Then we countinue clockwise
	// on the left side until we reach the starting point
	result := []Point{p1_points[upper_i], p2_points[upper_j]}
	left := p1_points[upper_i]
	left_index := 0
	right := p2_points[upper_j]
	right_index := 1

	for upper_j != lower_j {
		upper_j = (upper_j + 1) % len(p2_points)
		result = append(result, p2_points[upper_j])

		if p2_points[upper_j].X > right.X {
			right_index = len(result) - 1
			right = p2_points[upper_j]
		}
	}

	for lower_i != upper_i {
		result = append(result, p1_points[lower_i])

		if p1_points[lower_i].X < left.X {
			left_index = len(result) - 1
			left = p1_points[lower_i]
		}

		lower_i = (lower_i + 1) % len(p1_points)		
	}

	poly := Polygon{left_index, right_index, result}

	return poly
}

// Given two points, this method calculates the Y value of a point 
// having value X lying on the line defined by p1 and p2
func y_value(p1, p2 Point, x float64) float64 {
	m, b := equation_from(p1, p2)
	res := m * x + b
	return res
}

// Calculates the equation of the line defined by p1 and p2
func equation_from(p1, p2 Point) (float64, float64) {
	m := (p2.Y - p1.Y) / (p2.X - p1.X)
	b := p1.Y - p1.X * m

	return m, b
}
