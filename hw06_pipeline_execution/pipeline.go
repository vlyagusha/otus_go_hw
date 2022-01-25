package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)
	defer close(out)

	switch len(stages) {
	case 0:
		return out
	case 1:
		return stages[0](in)
	default:
		return ExecutePipeline(stages[0](in), done, stages[1:]...)
	}
}
