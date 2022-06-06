package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type nodeService struct {
	nodeRepo repository.Node
}

func (n nodeService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Node, int64) {
	if ownerReference == "" {
		return n.nodeRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return n.nodeRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

// NewNodeService returns service.Node type service
func NewNodeService(nodeRepo repository.Node) service.Node {
	return &nodeService{
		nodeRepo: nodeRepo,
	}
}
