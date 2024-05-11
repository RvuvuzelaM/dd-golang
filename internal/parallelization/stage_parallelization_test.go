package parallelization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelization(t *testing.T) {
	t.Run("everything can be done in parallel when there are no dependencies", func(t *testing.T) {
		stage1 := NewStage("Stage1")
		stage2 := NewStage("Stage2")

		stageParallelization := &StageParallelization{}
		sortedStages := stageParallelization.Of([]Stage{*stage1, *stage2})

		assert.Equal(t, 1, len(sortedStages.all))
	})

	t.Run("test simple dependencies", func(t *testing.T) {
		stage1 := NewStage("Stage1")
		stage2 := NewStage("Stage2")
		stage3 := NewStage("Stage3")
		stage4 := NewStage("Stage4")
		stage1.DependsOn(stage2)
		stage3.DependsOn(stage1)
		stage4.DependsOn(stage2)

		stageParallelization := &StageParallelization{}
		sortedStages := stageParallelization.Of([]Stage{*stage1, *stage2, *stage3, *stage4})

		assert.Equal(t, "Stage1 | Stage2, Stage3 | Stage4", sortedStages.Print())
	})

	t.Run("could not with cycle dependencies", func(t *testing.T) {
		stage1 := NewStage("Stage1")
		stage2 := NewStage("Stage2")
		stage1.DependsOn(stage2)
		stage2.DependsOn(stage1)

		stageParallelization := &StageParallelization{}
		sortedStages := stageParallelization.Of([]Stage{*stage1, *stage2})

		assert.Equal(t, 0, len(sortedStages.all))
	})
}
