package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		fmt.Println("Usage: convex-hull-go input output")
	} else {
		hull := read_file(arguments[0])
		write_file(hull, arguments[1])
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
		return Calculate_hull(points)
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
