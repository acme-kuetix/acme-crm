package transitions

import (
	"github.com/kuetix/engine/engine/domain/interfaces"
	"github.com/kuetix/engine/engine/workflow"
)

type crmTransitions struct {
	workflow.BaseServiceTransition
}

func NewCrmTransitions() interfaces.ServiceTransitions {
	return &crmTransitions{}
}
