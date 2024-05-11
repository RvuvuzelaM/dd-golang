package parallelization

import (
	"strings"
)

type ParallelStagesList struct {
	all []ParallelStages
}

func NewParallelStagesList(stages ...ParallelStages) ParallelStagesList {
	return ParallelStagesList{all: stages}
}

func EmptyParallelStagesList() ParallelStagesList {
	return NewParallelStagesList()
}

func (psl *ParallelStagesList) Add(newParallelStages ParallelStages) *ParallelStagesList {
	psl.all = append(psl.all, newParallelStages)
	return psl
}

func (psl *ParallelStagesList) Print() string {
	var prints []string
	for _, stages := range psl.all {
		prints = append(prints, stages.Print())
	}
	return strings.Join(prints, " | ")
}
