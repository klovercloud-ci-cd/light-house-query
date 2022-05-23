package logic

import (
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
)

type certificateService struct {
	certificateRepo repository.Certificate
}

func (c certificateService) Get(agent string, option v1.ResourceQueryOption) ([]v1.Certificate, int64) {
	return c.certificateRepo.Get(agent, option)
}

// NewCertificateService returns service.Certificate type service
func NewCertificateService(certificateRepo repository.Certificate) service.Certificate {
	return &certificateService{
		certificateRepo: certificateRepo,
	}
}
