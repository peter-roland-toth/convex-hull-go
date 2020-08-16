package main

import (
	"testing"
	// "fmt"
)

func TestMergeWithNullPolygons(t *testing.T) {
	p1 := []Point{}
	p2 := []Point{}

	result := merge(Polygon{0, 0, p1}, Polygon{0, 0, p2})

	if len(result.points) != 0 {
		t.Errorf("Result should be null polygon")
	}
}

func TestMergeWithTwoPoints(t *testing.T) {
	left_point := Point{-1.0, 0.0}
	right_point := Point{1.0, 0.0}

	p1 := Polygon{0, 0, []Point{left_point}}
	p2 := Polygon{0, 0, []Point{right_point}}

	result := merge(p1, p2)

	if len(result.points) != 2 {
		t.Errorf("Polygon should have two points")
	}

	if result.left_index != 0 {
		t.Errorf("Wrong left index")
	}
	
	if result.right_index != 1 {
		t.Errorf("Wrong right index")
	}
}

func TestMergeFourPoints(t *testing.T) {
	left_1 := Point{104.32917758136055, 139.67264895774775} 
	left_2 := Point{125.85522810221978, 72.58780155542837}
	right_1 := Point{135.0070591614026, 93.81656585134968} 
	right_2 := Point{157.8799203366949, 69.65367939960399}

	left := Polygon{1, 0, []Point{left_2, left_1}}
	right := Polygon{0, 1, []Point{right_1, right_2}}

	result := merge(left, right)

	if len(result.points) != 3 {
		t.Errorf("Polygon should be a triangle")
	}

	expected := []Point{left_1, right_2, left_2}

	for i := 0; i < len(result.points); i++ {
		if result.points[i] != expected[i] {
			t.Errorf("%v should match %v", result, expected)
			break
		}
	}

	if result.left_index != 0 {
		t.Errorf("Wrong left index")
	}

	if result.right_index != 1 {
		t.Errorf("Wrong right index")
	}
}

/*
func TestMergePointWithLine(t *testing.T) {
	left := Point{-1.0, 0.0}
	right_1 := Point{1.0, 1.0}
	right_2 := Point{1.0, -1.0}

	p1 := []Point{left}
	p2 := []Point{right_1, right_2}

	result := merge(p1, p2)
	expected := []Point{left, right_1, right_2}

	if len(result) != 3 {
		t.Errorf("Polygon should have three points")
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("%v should match %v", result, expected)
			break
		}
	}
}

func TestMergeLineWithPoint(t *testing.T) {
	left_1 := Point{-1.0, 1.0}
	left_2 := Point{-1.0, -1.0}
	right := Point{1.0, 1.0}
	
	p1 := []Point{left_1, left_2}
	p2 := []Point{right}

	result := merge(p1, p2)
	expected := []Point{left_1, right, left_2}

	if len(result) != 3 {
		t.Errorf("Polygon should have three points")
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("%v should match %v", result, expected)
			break
		}
	}
}

func TestMergeTriangleWithTriangle(t *testing.T) {
	left_1 := Point{-2.0, 0.0}
	left_2 := Point{-1.0, 1.0}
	left_3 := Point{-1.0, -1.0}
	right_1 := Point{1.0, 1.0}
	right_2 := Point{2.0, 0.0}
	right_3 := Point{1.0, -1.0}
	
	p1 := []Point{left_1, left_2, left_3}
	p2 := []Point{right_1, right_2, right_3}

	result := merge(p1, p2)
	expected := []Point{left_2, right_1, right_2, right_3, left_3, left_1}

	if len(result) != 6 {
		t.Errorf("Polygon should have six points")
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("%v should match %v", result, expected)
			break
		}
	}
}

func TestMergeTwoVerticalLines(t *testing.T) {
	left_1 := Point{-1.0, 0.0}
	left_2 := Point{-1.0, 1.0}
	left_3 := Point{-1.0, -1.0}
	right_1 := Point{1.0, 1.0}
	right_2 := Point{1.0, 0.0}
	right_3 := Point{1.0, -1.0}
	
	p1 := []Point{left_3, left_1, left_2}
	p2 := []Point{right_1, right_2, right_3}

	result := merge(p1, p2)

	if len(result) != 6 {
		t.Errorf("Polygon should have six points")
	}
}
*/

func TestMergingMultiplePoints(t *testing.T) {
	left_1 := Point{-3.0, 5.0}
	left_2 := Point{-2.0, 8.0}
	left_3 := Point{-1.0, 6.5}
	left_4 := Point{-2.2, 4.0}
	right_1 := Point{1.0, 4.5}
	right_2 := Point{2.0, 7.8}
	right_3 := Point{8.0, 4.5}
	right_4 := Point{2.0, -10.0}

	p1 := Polygon{0, 2, []Point{left_1, left_2, left_3, left_4}}
	p2 := Polygon{0, 2, []Point{right_1, right_2, right_3, right_4}}

	result := merge(p1, p2)

	if len(result.points) != 5 {
		t.Errorf("Polygon should have five points")
	}

	expected := []Point{left_2, right_2, right_3, right_4, left_1}

	for i := 0; i < len(result.points); i++ {
		if result.points[i] != expected[i] {
			t.Errorf("%v should match %v", result, expected)
			break
		}
	}

	if result.left_index != 4 {
		t.Errorf("Wrong left index %v", result.left_index)
	}

	if result.right_index != 2 {
		t.Errorf("Wrong right index %v", result.right_index)
	}
}

/*
func TestMergingMultiplePoints2(t *testing.T) {
	left_1 := Point{8, 12}
	left_2 := Point{10, 18}
	right_1 := Point{12, 11.1}
	right_2 := Point{12, 21.3}
	right_3 := Point{15, 2.5}

	p1 := []Point{left_1, left_2}
	p2 := []Point{right_1, right_2, right_3}

	result := merge(p1, p2)

	if len(result) != 4 {
		t.Errorf("Polygon should have four points")
	}
}

func TestMergeWithPointsOfRectangle(t *testing.T) {
	upper_left := Point{-1.0, 1.0}
	lower_left := Point{-1.0, -1.0}
	upper_right := Point{1.0, 1.0}
	lower_right := Point{1.0, -1.0}

	p1 := []Point{upper_left, lower_left}
	p2 := []Point{upper_right, lower_right}

	result := merge(p1, p2)
	// todo: check actual array content
	if len(result) != 4 {
		t.Errorf("Polygon should have four points")
	}
}

func TestPartitionOneElement(t *testing.T) {
	e1 := Point{2, 10}

	p := []Point{e1}
	l, r := partition(p)

	if len(l) != 0 && len(r) != 1 {
		t.Errorf("Partitioning one element is wrong")
	}
}

func TestPartitionTwoElements(t *testing.T) {
	e1 := Point{8, 12}
	e2 := Point{10, 18}

	p := []Point{e1, e2}
	l, r := partition(p)

	if len(l) != 1 && len(r) != 1 {
		t.Errorf("Partitioning two elements is wrong")
	}
}

func TestPartitionEmptyArray(t *testing.T) {
	p := []Point{}
	l, r := partition(p)

	if len(l) != 0 && len(r) != 0 {
		t.Errorf("Partitioning zero element0 is wrong")
	}
}

func TestAreaOfRectangle(t *testing.T) {
	e1 := Point{-1.0, 1.0}
	e2 := Point{2.0, 1.0}
	e3 := Point{2.0, -1.0}
	e4 := Point{-1.0, -1.0}

	p := []Point{e1, e2, e3, e4}

	area := area(p)

	if area != 6.0 {
		t.Errorf("Wrong area %f. Should be 6.0", area)
	}
}
*/
