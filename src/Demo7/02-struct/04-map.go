package main

import "fmt"

type Vertex2 struct {
	Lat, long float64
}

var (
	m = map[string]Vertex2{
		"google": {
			Lat:  12,
			long: 23,
		},
		"youtube": {
			Lat:  122,
			long: 222,
		},
	}
)

func main() {
	fmt.Println(m)
}
