package main

import (
	"math"
)

type Node struct {
	name              string
	size              int
	total             int
	branches          map[float64]*Node
	attr_values       Counter
	attr_class_values map[float64]Counter
}

func (node *Node) Test(samples []Sample) (int, int) {
	var (
		hits     = 0
		failures = 0
	)
	for _, s := range samples {
		p := node.predict(s)
		if p == nil {
			failures += 1

			continue
		}
		value := s.GetValue("status")
		if p.Best() == value {
			hits += 1
		} else {
			failures += 1
		}
	}
	return hits, failures
}

func (node *Node) SetLeaf(value float64, dist *Distribution) {
	_, ok := node.attr_values[value]
	if !ok {
		node.attr_values[value] = 0.0
	}
	node.attr_values[value] += 1.0
	node.total += 1
	_, ok = node.attr_class_values[value]
	if !ok {
		node.attr_class_values[value] = make(Counter)
	}
	for k, v := range dist.Counter {
		_, ok := node.attr_class_values[value][k]
		if !ok {
			node.attr_class_values[value][k] = 0.0
		}
		node.attr_class_values[value][k] += v
	}
}

func (node *Node) get_value_dist(value float64) *Distribution {
	status := "new"
	if value == StatusDegraded {
		status = "degraded"
	}

	dist := NewDistribution(status)
	for k, v := range node.attr_class_values[value] {
		dist.Add(k, v)
	}
	return dist
}

func (node *Node) get_values() []float64 {
	temp := make(map[float64]bool)
	for k, _ := range node.attr_values {
		temp[k] = true
	}
	for k, _ := range node.branches {
		temp[k] = true
	}
	values := make([]float64, 0)
	for value, _ := range temp {
		values = append(values, value)
	}
	return values
}

func (node *Node) get_attr_value_from_sample(sample Sample) (float64, bool) {
	value := sample.GetValue(node.name)
	values := node.get_values()
	var (
		n          float64 = math.MaxFloat64
		nearest    float64 = 0.0
		hasNearest bool    = false
	)
	for _, v := range values {
		if value == v {
			return value, true
		}
		if n > math.Abs(v-value) {
			n = math.Abs(v - value)
			nearest = v
			hasNearest = true
		}
	}
	return nearest, hasNearest
}

func (node *Node) predict(sample Sample) *Distribution {
	value, ok := node.get_attr_value_from_sample(sample)
	if !ok {
		return nil
	}
	for v, branch := range node.branches {
		if v == value {
			return branch.predict(sample)
		}
	}
	return node.get_value_dist(value)
}

func (node *Node) create_branch(val float64, subtree *Node) {
	node.branches[val] = subtree
}

func NewNode(name string, size int) *Node {
	return &Node{
		name:              name,
		size:              size,
		branches:          make(map[float64]*Node),
		total:             0,
		attr_values:       make(Counter),
		attr_class_values: make(map[float64]Counter),
	}
}
