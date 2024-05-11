package parallelization

type StageParallelization struct{}

func (sp *StageParallelization) Of(stages []Stage) ParallelStagesList {
	for _, stage := range stages {
		if hasCycle(&stage, make(map[*Stage]bool), make(map[*Stage]bool)) {
			return EmptyParallelStagesList()
		}
	}

	return sortNodes(stages, EmptyParallelStagesList())
}

func hasCycle(node *Stage, visited, recStack map[*Stage]bool) bool {
	if recStack[node] {
		return true
	}
	if visited[node] {
		return false
	}
	visited[node] = true
	recStack[node] = true
	for dep := range node.dependencies {
		if hasCycle(dep, visited, recStack) {
			return true
		}
	}
	recStack[node] = false
	return false
}

func sortNodes(remainingNodes []Stage, processedParallelNodes ParallelStagesList) ParallelStagesList {
	nodesWithoutDependencies := withAllDependenciesPresentIn(
		remainingNodes,
		alreadyProcessedNodes(processedParallelNodes),
	)

	if len(nodesWithoutDependencies) == 0 {
		return processedParallelNodes
	}

	newprocessedParallelNodes := processedParallelNodes.Add(NewParallelStages(nodesWithoutDependencies))
	newremainingNodes := []Stage{}
	for _, stage := range remainingNodes {
		if !contains(nodesWithoutDependencies, stage) {
			newremainingNodes = append(newremainingNodes, stage)
		}
	}
	return sortNodes(newremainingNodes, *newprocessedParallelNodes)
}

func alreadyProcessedNodes(processedParallelNodes ParallelStagesList) []Stage {
	alreadyProcessedNodes := []Stage{}
	for _, node := range processedParallelNodes.all {
		for _, stage := range node.stages {
			alreadyProcessedNodes = append(alreadyProcessedNodes, stage)
		}
	}
	return alreadyProcessedNodes
}

func withAllDependenciesPresentIn(toCheck, presentIn []Stage) []Stage {
	var result []Stage
	for _, node := range toCheck {
		if allDependenciesPresent(node, presentIn) {
			result = append(result, node)
		}
	}
	return result
}

func allDependenciesPresent(node Stage, presentIn []Stage) bool {
	for dep := range node.dependencies {
		if !contains(presentIn, *dep) {
			return false
		}
	}
	return true
}

func contains(slice []Stage, value Stage) bool {
	for _, item := range slice {
		if item.Equals(value) {
			return true
		}
	}
	return false
}
