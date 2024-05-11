package parallelization

import (
	"sort"
	"strings"
)

// ParallelStages represents a group of stages that are processed in parallel
type ParallelStages struct {
	stages []Stage
}

// NewParallelStages creates a new ParallelStages instance
func NewParallelStages(stages []Stage) ParallelStages {
	return ParallelStages{stages: stages}
}

// Print outputs a comma-separated, sorted list of stage names
func (ps *ParallelStages) Print() string {
	var names []string
	for _, stage := range ps.stages {
		names = append(names, stage.Name())
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}
