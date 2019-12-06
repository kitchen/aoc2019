package day6

import (
	"log"
	"strings"

	"github.com/goombaio/dag"
)

type OrbitalDAG struct {
	*dag.DAG
}

func NewOrbitalDag() *OrbitalDAG {
	return &OrbitalDAG{DAG: dag.NewDAG()}
}

// parent)child
func (graph *OrbitalDAG) AddOrbit(orbit string) error {
	bodies := strings.Split(orbit, ")")
	bodyID := bodies[0]
	satelliteID := bodies[1]

	body, err := graph.GetVertex(bodyID)
	if err != nil {
		body = dag.NewVertex(bodyID, bodyID)
		graph.AddVertex(body)
	}

	satellite, err := graph.GetVertex(satelliteID)
	if err != nil {
		satellite = dag.NewVertex(satelliteID, satelliteID)
		graph.AddVertex(satellite)
	}

	return graph.AddEdge(body, satellite)
}

func (graph *OrbitalDAG) Ancestors(vertex *dag.Vertex) ([]*dag.Vertex, error) {
	ancestors := []*dag.Vertex{}
	parents, err := graph.Predecessors(vertex)
	if err != nil {
		return nil, err
	}

	for _, parent := range parents {
		predecessors, err := graph.Ancestors(parent)
		if err != nil {
			return nil, err
		}
		ancestors = append(ancestors, parent)
		ancestors = append(ancestors, predecessors...)
	}
	return ancestors, nil
}

func (graph *OrbitalDAG) Distances(vertex *dag.Vertex, distance int) int {
	distances := 0
	children, err := graph.Successors(vertex)
	if err != nil {
		log.Fatal("foo")
	}
	for _, child := range children {
		distances += distance
		distances += graph.Distances(child, distance+1)
	}
	return distances
}
