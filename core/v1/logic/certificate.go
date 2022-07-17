package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type certificateService struct {
	certificateRepo repository.Certificate
}

func (c certificateService) GetById(id, agent, processId string) v1.Certificate {
	return c.certificateRepo.GetById(id, agent, processId)
}

func (c certificateService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Certificate, int64) {
	if ownerReference == "" {
		return c.certificateRepo.GetByAgentAndProcessId(agent, processId, option)
	}
	return c.certificateRepo.GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId, option)
}

func (c certificateService) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Certificate {
	return c.certificateRepo.GetByAgentAndProcessIdWithoutPagination(agent, processId)
}

// NewCertificateService returns service.Certificate type service
func NewCertificateService(certificateRepo repository.Certificate) service.Certificate {
	return &certificateService{
		certificateRepo: certificateRepo,
	}
}
