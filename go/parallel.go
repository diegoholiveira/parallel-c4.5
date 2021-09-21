package main

import (
	"math"
	"sync"
)

func parallel(cpu int) TreeGenerator {
	chooser := choose_attribute_async(cpu)

	return func(samples []Sample) *Node {
		c45 := &C45{
			ChooseAttribute: chooser,
		}
		return c45.generate(samples, DataAttributes, DataClassAttr)
	}
}

type (
	job struct {
		dataset    []Sample
		attr       string
		class_attr string
	}
	output struct {
		value float64
		attr  string
	}
)

func worker(wg *sync.WaitGroup, jobs <-chan job, results chan<- output) {
	for j := range jobs {
		v := gain(j.dataset, j.attr, j.class_attr)
		results <- output{
			attr:  j.attr,
			value: v,
		}
		wg.Done()
	}
}

func choose_attribute_async(cpu int) Chooser {
	return func(dataset []Sample, attributes []string, class_attr string) (bool, string, []float64) {
		jobs := make(chan job, len(attributes)-1)
		for _, attr := range attributes {
			if attr == class_attr {
				continue
			}

			jobs <- job{
				dataset:    dataset,
				attr:       attr,
				class_attr: class_attr,
			}
		}
		close(jobs)

		results := make(chan output, len(attributes)-1)
		var wg sync.WaitGroup
		wg.Add(len(attributes) - 1)
		for i := 0; i < cpu; i++ {
			go worker(&wg, jobs, results)
		}

		wg.Wait()
		close(results)

		var (
			best    string  = ""
			value   float64 = math.MaxFloat64 * -1
			hasBest bool    = false
		)

		for result := range results {
			if result.value > value {
				value = result.value
				best = result.attr
				hasBest = true
			}
		}

		values := make([]float64, 0)
		if hasBest {
			values = unique_values(dataset, best)
		}
		return hasBest, best, values
	}
}
