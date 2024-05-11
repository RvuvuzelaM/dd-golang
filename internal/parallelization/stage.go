package parallelization

type Stage struct {
	dependencies map[*Stage]struct{}
	stageName    string
}

type ResourceName struct {
	Name string
}

// NewStage initializes a new Stage with optional parameters
func NewStage(stageName string) Stage {
	stage := Stage{
		dependencies: make(map[*Stage]struct{}),
		stageName:    stageName,
	}
	return stage
}

func (s Stage) Name() string {
	return s.stageName
}

func (s Stage) DependsOn(stage *Stage) Stage {
	s.dependencies[stage] = struct{}{}
	return s
}

func (s Stage) GetDependencies() []*Stage {
	dep := []*Stage{}
	for k := range s.dependencies {
		dep = append(dep, k)
	}
	return dep
}

func (s Stage) Equals(stage Stage) bool {
	return s.stageName == stage.stageName
}
