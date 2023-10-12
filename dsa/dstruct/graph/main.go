package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	data     int
	adjacent []*Vertex
}

//add vertex
func (g *Graph) addVertex(v *Vertex) {
	//check for duplicate
	if contains(g.vertices, v.data) {
		return
	}
	g.vertices = append(g.vertices, v)
}

//add addjacent
func (v *Vertex) addAdjacent(to *Vertex) {
	//check for duplicate
	if contains(v.adjacent, to.data) {
		return
	}
	v.adjacent = append(v.adjacent, to)
}

//contains
func contains(v []*Vertex, d int) bool {
	for _, v1 := range v {
		if v1.data == d {
			return true
		}
	}
	return false
}

//get vertex
func (g *Graph) getVertex(d int) (v *Vertex) {

	for _, v := range g.vertices {
		if v.data == d {
			return v
		}
	}
	return nil
}

//delete
func (v *Vertex) deleteAdjacent(d int) {
	for i, v1 := range v.adjacent {
		if v1.data == d {
			v.adjacent = append(v.adjacent[:i], v.adjacent[i+1:]...)
		}
	}

}

//delete
func (g *Graph) deleteVetex(d int) {
	for i, v := range g.vertices {
		if v.data == d {
			g.vertices = append(g.vertices[:i], g.vertices[i+1:]...)
		}
		v.deleteAdjacent(d)
	}
}

func (g *Graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("vertex %d : ", v.data)
		for _, v1 := range v.adjacent {
			fmt.Printf("%d", v1.data)
			for _, v2 := range v1.adjacent {
				fmt.Printf("[%d],", v2.data)
			}
		}
		fmt.Println()
	}
}
func main() {
	g := &Graph{}
	v1 := &Vertex{data: 10}
	g.addVertex(v1)
	v1 = &Vertex{data: 35}
	g.addVertex(v1)
	v1 = &Vertex{data: 26}
	g.addVertex(v1)
	v1 = &Vertex{data: 26} //duplicate vertex
	g.addVertex(v1)
	//adding adjacent for 10
	if v1 = g.getVertex(10); v1 != nil {
		v2 := g.getVertex(35)
		if v2 != nil {
			v1.addAdjacent(v2)
		}
		if v2 = g.getVertex(26); v2 != nil {
			v1.addAdjacent(v2)
		}
		//duplicate adjacent
		if v2 = g.getVertex(35); v2 != nil {
			v1.addAdjacent(v2)
		}
	}
	//adding adjacent for 35
	if v1 = g.getVertex(35); v1 != nil {
		v2 := g.getVertex(26)
		if v2 != nil {
			v1.addAdjacent(v2)
		}
		//donst exists
		v2 = g.getVertex(56)
		if v2 != nil {
			v1.addAdjacent(v2)
		}
	}
	v1 = g.getVertex(10)
	v1.deleteAdjacent(26)
	g.deleteVetex(26) //also update adjacents
	g.print()
}
