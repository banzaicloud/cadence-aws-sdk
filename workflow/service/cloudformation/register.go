package cloudformation

import (
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
)

const activityNamePrefix = "aws-cloudformation-"

// Register registers the activities in a worker.
func (a Activities) Register(worker worker.Worker) {
	worker.RegisterActivityWithOptions(a, activity.RegisterOptions{Name: activityNamePrefix})
}
