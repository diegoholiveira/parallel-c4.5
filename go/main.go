package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type (
	Stat struct {
		Method    string `json:"method"`
		Execution int    `json:"execution"`
		Elapsed   string `json:"elapsed"`
		Hits      int    `json:"hits"`
		Failures  int    `json:"failures"`
	}

	TreeGenerator func([]Sample) *Node
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var (
		dataDir         = filepath.Join(dir, "..", "data")
		datasetFilename = filepath.Join(dataDir, "ts-large.csv")
		samplesFilename = filepath.Join(dataDir, "samples.csv")
	)

	dataset := fromCSV(datasetFilename)
	samples := fromCSV(samplesFilename)

	stats := make([]Stat, 0)
	for i := 0; 1 > i; i++ {
		stats = append(stats, get_stats_from_execution(
			"sequential",
			sequential,
			dataset,
			samples,
			i,
		))
		stats = append(stats, get_stats_from_execution(
			"parallel_2",
			parallel(2),
			dataset,
			samples,
			i,
		))
		stats = append(stats, get_stats_from_execution(
			"parallel_4",
			parallel(4),
			dataset,
			samples,
			i,
		))
		stats = append(stats, get_stats_from_execution(
			"parallel_8",
			parallel(8),
			dataset,
			samples,
			i,
		))
		stats = append(stats, get_stats_from_execution(
			"parallel_16",
			parallel(16),
			dataset,
			samples,
			i,
		))
	}

	file, _ := json.MarshalIndent(stats, "", "    ")

	output := "stats_" + time.Now().Format("20060102150405") + ".json"
	_ = ioutil.WriteFile(output, file, 0644)
}

func get_stats_from_execution(method string, f TreeGenerator, dataset, samples []Sample, i int) Stat {
	start := time.Now()
	tree := f(dataset)
	elapsed := time.Since(start)

	hits, failures := tree.Test(samples)

	s := Stat{
		Method:    method,
		Execution: i + 1,
		Elapsed:   durationToString(elapsed),
		Hits:      hits,
		Failures:  failures,
	}
	fmt.Println(s)
	return s
}

func durationToString(t time.Duration) string {
	var (
		hours   int64 = 0
		minutes int64 = 0
		seconds int64 = int64(t.Seconds())
	)
	if seconds > 59 {
		hours = seconds / int64(60*60)
		seconds = seconds % int64(60*60)
		minutes = seconds / int64(60)
		seconds = seconds % int64(60)
	}
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
