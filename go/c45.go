package main

type (
	Chooser func([]Sample, []string, string) (bool, string, []float64)

	C45 struct {
		ChooseAttribute Chooser
	}
)

func (c45 *C45) generate(samples []Sample, attributes []string, class_attr string) *Node {
	ok, best, values := c45.ChooseAttribute(samples, attributes, class_attr)
	if !ok {
		return nil
	}

	node := NewNode(best, len(samples))

	for _, value := range values {
		sub := create_subdataset(samples, best, value)
		dist := create_distribution(class_attr, sub)
		if is_leaf(sub, attributes, dist) {
			node.SetLeaf(value, dist)

			continue
		}

		sub_attr := make([]string, 0)
		for _, a := range attributes {
			if a == best {
				continue
			}
			sub_attr = append(sub_attr, a)
		}

		child := c45.generate(sub, sub_attr, class_attr)
		if child != nil {
			node.create_branch(value, child)
		}
	}
	return node
}

func unique_values(dataset []Sample, attr string) []float64 {
	tmp := make(map[float64]bool)
	for _, s := range dataset {
		tmp[s.GetValue(attr)] = true
	}
	values := make([]float64, 0)
	for v, _ := range tmp {
		values = append(values, v)
	}
	return values
}

func create_subdataset(dataset []Sample, best string, value float64) []Sample {
	s := make([]Sample, 0)
	for _, sample := range dataset {
		if sample.GetValue(best) == value {
			s = append(s, sample)
		}
	}
	return s
}

func is_leaf(dataset []Sample, attributes []string, dist *Distribution) bool {
	return len(dataset) == 0 ||
		0 >= len(attributes)-1 ||
		1 >= len(dist.Counter)
}
