package v1

import (
	"github.com/go-bongo/bongo"
)

type k8sClusterRoleBinding struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Subjects   []Subject `json:"subjects,omitempty" protobuf:"bytes,2,rep,name=subjects" bson:"subjects"`
	RoleRef    RoleRef   `json:"roleRef" protobuf:"bytes,3,opt,name=roleRef" bson:"roleRef"`
}

type ClusterRoleBinding struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                k8sClusterRoleBinding `bson:"obj" json:"obj"`
	KubeClusterId      string                `bson:"kubeClusterId" json:"kubeClusterId"`
}
