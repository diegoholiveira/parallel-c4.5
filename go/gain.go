package main

import "math"

type Counter map[float64]float64

func frequency(dataset []Sample, attr string) Counter {
	counter := make(Counter)
	for _, s := range dataset {
		value := s.GetValue(attr)

		v, ok := counter[value]
		if !ok {
			counter[value] = 1.0

			continue
		}

		counter[value] = v + 1.0
	}
	return counter
}

func logN(x float64, n float64) float64 {
	return math.Log(x) / math.Log(n)
}

func entropy(dataset []Sample, attr string) float64 {
	counter := frequency(dataset, attr)

	var (
		n int     = 2
		s float64 = 0
		e float64 = 0
	)

	if len(dataset) > n {
		n = len(dataset)
	}

	for _, value := range counter {
		s += value
	}

	for _, value := range counter {
		e += (value / s) * logN(value/s, float64(n))
	}

	return e * -1
}

func getSubset(dataset []Sample, attr string, value float64) []Sample {
	s := make([]Sample, 0)
	for _, sample := range dataset {
		if sample.GetValue(attr) == value {
			s = append(s, sample)
		}
	}
	return s
}

func gain(dataset []Sample, attr, classAttr string) float64 {
	var (
		s float64 = 0.0
		e float64 = 0.0
	)

	values := frequency(dataset, attr)

	for _, value := range values {
		s += value
	}

	for key, value := range values {
		prob := value / s
		subset := getSubset(dataset, attr, key)
		e += prob * entropy(subset, classAttr)
	}
	return entropy(dataset, classAttr) - e
}
