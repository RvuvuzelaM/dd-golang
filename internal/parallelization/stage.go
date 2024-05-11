package parallelization

type Stage struct {
	dependencies map[*Stage]struct{}
	stageName    string
}

type ResourceName struct {
	Name string
}

// NewStage initializes a new Stage with optional parameters
func NewStage(stageName string, opts ...func(*Stage)) *Stage {
	stage := &Stage{
		dependencies: make(map[*Stage]struct{}),
		stageName:    stageName,
	}
	for _, opt := range opts {
		opt(stage)
	}
	return stage
}

// WithDependencies adds dependencies to a Stage
func WithDependencies(dependencies ...*Stage) func(*Stage) {
	return func(s *Stage) {
		for _, dep := range dependencies {
			s.dependencies[dep] = struct{}{}
		}
	}
}

func (s *Stage) Name() string {
	return s.stageName
}

func (s *Stage) DependsOn(stage *Stage) *Stage {
	s.dependencies[stage] = struct{}{}
	return s
}

func (s *Stage) Equals(stage *Stage) bool {
	return s.stageName == stage.stageName
}
