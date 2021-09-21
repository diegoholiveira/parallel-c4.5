package main

type Distribution struct {
	Counter Counter
	Total   uint64
	Attr    string
}

func (d *Distribution) Add(key, value float64) {
	v, ok := d.Counter[key]
	if !ok {
		d.Counter[key] = value
	} else {
		d.Counter[key] = v + value
	}

	d.Total += uint64(value)
}

func (d *Distribution) Best() float64 {
	var (
		max  float64
		freq float64
	)

	for key, value := range d.Counter {
		if value > max {
			max = value
			freq = key
		}
	}

	return freq
}

func NewDistribution(attr string) *Distribution {
	return &Distribution{
		Counter: make(Counter),
		Total:   0,
		Attr:    attr,
	}
}

func create_distribution(class_attr string, dataset []Sample) *Distribution {
	dist := NewDistribution(class_attr)
	for _, s := range dataset {
		dist.Add(s.GetValue(class_attr), 1.0)
	}
	return dist
}
