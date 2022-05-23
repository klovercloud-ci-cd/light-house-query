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

// GetV1NodeService returns service.Node
func GetV1NodeService() service.Node {
	return logic.NewNodeService(mongo.NewNodeRepository(3000))
}

// GetV1ClusterRoleService returns service.ClusterRole
func GetV1ClusterRoleService() service.ClusterRole {
	return logic.NewClusterRoleService(mongo.NewClusterRoleRepository(3000))
}

// GetV1ClusterRoleBindingService returns service.ClusterRoleBinding
func GetV1ClusterRoleBindingService() service.ClusterRoleBinding {
	return logic.NewClusterRoleBindingService(mongo.NewClusterRoleBindingRepository(3000))
}

// GetV1ConfigMapService returns service.ConfigMap
func GetV1ConfigMapService() service.ConfigMap {
	return logic.NewConfigMapService(mongo.NewConfigMapRepository(3000))
}

// GetV1DaemonSetService returns service.DaemonSet
func GetV1DaemonSetService() service.DaemonSet {
	return logic.NewDaemonSetService(mongo.NewDaemonSetRepository(3000))
}

// GetV1DeploymentService returns service.Deployment
func GetV1DeploymentService() service.Deployment {
	return logic.NewDeploymentService(mongo.NewDeploymentRepository(3000))
}

// GetV1PersistentVolumeService returns service.PersistentVolume
func GetV1PersistentVolumeService() service.PersistentVolume {
	return logic.NewPersistentVolumeService(mongo.NewPersistentVolumeRepository(3000))
}

// GetV1PersistentVolumeClaimService returns service.PersistentVolumeClaim
func GetV1PersistentVolumeClaimService() service.PersistentVolumeClaim {
	return logic.NewPersistentVolumeClaimService(mongo.NewPersistentVolumeClaimRepository(3000))
}
