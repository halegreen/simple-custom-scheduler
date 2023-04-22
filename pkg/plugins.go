package plugins

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	Name = "simple"
)

type SimplePlugin struct {
}

// NewSimplePlugin initializes and returns a new simple plugin.
func NewSimplePlugin(obj runtime.Object, f framework.Handle) (framework.Plugin, error) {
	return &SimplePlugin{}, nil
}

func (p *SimplePlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	return framework.NewStatus(framework.Success)
}

func (p *SimplePlugin) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	return framework.NewStatus(framework.Success)
}

func (cs *SimplePlugin) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

func (p *SimplePlugin) PostFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, filteredNodeStatusMap framework.NodeToStatusMap) (*framework.PostFilterResult, *framework.Status) {
	return nil, framework.NewStatus(framework.Success)
}

func (p *SimplePlugin) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	return 0, framework.NewStatus(framework.Success)
}

func (p *SimplePlugin) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

func (p *SimplePlugin) Reserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	return framework.NewStatus(framework.Success)
}

func (p *SimplePlugin) Unreserve(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {
}

func (p *SimplePlugin) Name() string {
	return Name
}
