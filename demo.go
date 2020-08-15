package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"math"
	"time"
	"bufio"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		read_standard_input()
	} else if len(arguments) == 2 {
		points := read_file(arguments[0])
		t1 := time.Now()
		hull := Calculate_hull(points)
		t2 := time.Now()
		fmt.Println("Number of points: ", len(points))
		diff := t2.Sub(t1)
		fmt.Println("Duration: ", diff)
		fmt.Println("Speed: ", 1000.0 * int64(len(points)) / diff.Microseconds(), "points/ms")
		fmt.Println(area(hull))
		write_file(hull, arguments[1])
	} else {
		fmt.Println("Usage: convex-hull-go input output")
	}
}

func read_standard_input() {
	var err error
	var points []Point

	for err == nil {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		err = json.Unmarshal([]byte(text), &points)

		if err == nil {
			write_file(points, fmt.Sprintf("%d.json", len(points)))
			hull := Calculate_hull(points)
			fmt.Println(area(hull))
		}	
	}
}

func read_file(name string) []Point {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
    defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var points []Point

	err = json.Unmarshal(byteValue, &points)

	if err == nil {
		return points
	}

	return []Point{}
}

func write_file(points []Point, name string) {					
	b, err := json.Marshal(points)

	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(name, b, 0644)

	if err != nil {
		fmt.Println(err)
	}
}

// Calculates the area of a polygon. It assumes that the points 
// are ordered clockwise or counter-clockwise
func area(p []Point) float64 {
	if len(p) <= 2 {
		return 0.0
	}

	sum := 0.0
	for i := 0; i < len(p); i++ {
		current := p[i]
		next := p[(i+1) % len(p)]
		sum += current.X * next.Y - current.Y * next.X
	}

	return math.Abs(sum / 2.0)
}
