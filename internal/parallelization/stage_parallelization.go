package parallelization

type StageParallelization struct{}

func (sp *StageParallelization) Of(stages []Stage) ParallelStagesList {
	for _, stage := range stages {
		if stage.HasCycle(make(map[*Stage]bool), make(map[*Stage]bool)) {
			return EmptyParallelStagesList()
		}
	}

	parallelStage := NewParallelStages(stages)
	parallelStages := NewParallelStagesList(parallelStage)
	return parallelStages
}
