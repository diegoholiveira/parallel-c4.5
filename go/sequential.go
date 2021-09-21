package main

import (
	"math"
)

func sequential(samples []Sample) *Node {
	c45 := &C45{
		ChooseAttribute: choose_attribute,
	}
	return c45.generate(samples, DataAttributes, DataClassAttr)
}

func choose_attribute(dataset []Sample, attributes []string, class_attr string) (bool, string, []float64) {
	var (
		best    string  = ""
		value   float64 = math.MaxFloat64 * -1
		hasBest bool    = false
	)

	for _, attr := range attributes {
		if attr == class_attr {
			continue
		}

		v := gain(dataset, attr, class_attr)
		if v > value {
			value = v
			best = attr
			hasBest = true
		}

	}

	values := make([]float64, 0)
	if hasBest {
		values = unique_values(dataset, best)
	}
	return hasBest, best, values
}
