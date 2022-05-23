package dependency

import (
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/logic"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/service"
	"github.com/klovercloud-ci-cd/light-house-query/repository/v1/mongo"
)

// GetV1PodService returns service.Pod
func GetV1PodService() service.Pod {
	return logic.NewPodService(mongo.NewPodRepository(3000))
}

// GetV1CertificateService returns service.Certificate
func GetV1CertificateService() service.Certificate {
	return logic.NewCertificateService(mongo.NewCertificateRepository(3000))
}
