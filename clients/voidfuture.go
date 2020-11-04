
package clients

import (
	"go.uber.org/cadence/workflow"
)

type VoidFuture struct {
	Future workflow.Future
}

func (r *VoidFuture) Get(ctx workflow.Context) error {
	return r.Future.Get(ctx, nil)
}

func NewVoidFuture(future workflow.Future) *VoidFuture {
	return &VoidFuture{Future: future}
}
