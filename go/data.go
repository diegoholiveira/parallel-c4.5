package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	TorquePosition     int = 1
	PCUTSpeedPosition  int = 4
	PSVOLSpeedPosition int = 6
	VAXSpeedPosition   int = 8
	ModePosition       int = 9
	StatusPosition     int = 10

	StatusNormal   float64 = 1.0
	StatusDegraded float64 = 2.0
)

type Sample struct {
	Torque     float64
	PCUTSpeed  float64
	PSVOLSpeed float64
	VAXSpeed   float64
	Mode       float64
	Status     string
}

func (s Sample) GetValue(attr string) float64 {
	switch attr {
	case "torque":
		return s.Torque
	case "pcut_speed":
		return s.PCUTSpeed
	case "psvol_speed":
		return s.PSVOLSpeed
	case "vax_speed":
		return s.VAXSpeed
	case "mode":
		return s.Mode
	default:
		if s.Status == "normal" {
			return StatusNormal
		}

		return StatusDegraded
	}
}

func parseSample(row []string) Sample {
	torque, _ := strconv.ParseFloat(row[TorquePosition], 64)
	pcutSpeed, _ := strconv.ParseFloat(row[PCUTSpeedPosition], 64)
	psvolSpeed, _ := strconv.ParseFloat(row[PSVOLSpeedPosition], 64)
	vaxSpeed, _ := strconv.ParseFloat(row[VAXSpeedPosition], 64)
	mode, _ := strconv.ParseFloat(row[ModePosition], 64)
	status := row[StatusPosition]

	return Sample{
		Torque:     torque,
		PCUTSpeed:  pcutSpeed,
		PSVOLSpeed: psvolSpeed,
		VAXSpeed:   vaxSpeed,
		Mode:       mode,
		Status:     status,
	}
}

func fromCSV(filename string) []Sample {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	samples := make([]Sample, 0)

	for {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		samples = append(samples, parseSample(record))
	}

	return samples
}
