package domain

import "time"

type Statistics struct {
	TotalTasks                 int
	CompletedTasks             int
	CompletedRateTasks         *float64
	AverageCompletionTimeTasks *time.Duration
}

func NewStatistics(
	totalTasks int,
	completedTasks int,
	completedRateTasks *float64,
	averageCompletionTimeTasks *time.Duration,
) Statistics {
	return Statistics{
		TotalTasks:                 totalTasks,
		CompletedTasks:             completedTasks,
		CompletedRateTasks:         completedRateTasks,
		AverageCompletionTimeTasks: averageCompletionTimeTasks,
	}
}
