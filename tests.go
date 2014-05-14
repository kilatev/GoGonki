package main

import (
    "testing"
    "fmt"
    )

func TestCarMoveTo(t *testing.T) {
    car := Car{}
    car.Point.X = 20
    car.Point.Y = 30
    point := car.MoveTo(Point{20, 20})
    if point.X == 20{
        fmt.Println("Yes")
    }
}