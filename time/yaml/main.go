package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Job struct {
	Start Date `yaml:"start_date"`
	End   Date `yaml:"end_date"`
}

type Date struct {
	time.Time
}

func (d *Date) Before(date Date) bool {
	return d.Time.Before(date.Time)
}

func (d *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	if err := unmarshal(&buf); err != nil {
		return nil
	}

	tt, err := time.Parse("2006-01-02", strings.TrimSpace(buf))
	if err != nil {
		return err
	}
	d.Time = tt
	return nil
}

func (d Date) MarshalYAML() (interface{}, error) {
	return d.Time.Format("2006-01-02"), nil
}

func main() {
	data, err := os.ReadFile("data.yaml")
	if err != nil {
		panic(err)
	}

	var job struct {
		Job Job `yaml:"job"`
	}

	err = yaml.Unmarshal(data, &job)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Start: %s - End: %s\n", job.Job.Start, job.Job.End)
	fmt.Printf("Star is minor than End: %t", job.Job.Start.Before((job.Job.End)))
}
