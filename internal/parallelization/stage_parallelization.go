package parallelization

type StageParallelization struct{}

func (sp *StageParallelization) Of(stages []Stage) *ParallelStagesList {
	return EmptyParallelStagesList()
}
