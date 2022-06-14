package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type agentService struct {
	agentRepo repository.Agent
}

func (a agentService) Get(companyId string) []v1.Agent {
	return a.agentRepo.Get(companyId)
}

// NewAgentService returns service.Agent type service
func NewAgentService(agentRepo repository.Agent) service.Agent {
	return &agentService{
		agentRepo: agentRepo,
	}
}
