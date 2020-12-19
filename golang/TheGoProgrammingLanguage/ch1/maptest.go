package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, 74.39967,
	}
	m["test"] = Vertex{
		12.0, 100,
	}
	m["hulk"] = Vertex{
		13.0, 100,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}
