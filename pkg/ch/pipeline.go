package ch

type PipelineStage[T any] func(done <-chan struct{}, in <-chan T) <-chan T
