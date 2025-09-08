package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		Lat:  43.342342,
		Long: 23.234266,
	}
	fmt.Println(m)

	delete(m, "Bell Labs")

	fmt.Println(m)
}
