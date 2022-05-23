package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sRoleBinding struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Subjects   []Subject `json:"subjects,omitempty" protobuf:"bytes,2,rep,name=subjects" bson:"subjects"`
	RoleRef    RoleRef   `json:"roleRef" protobuf:"bytes,3,opt,name=roleRef" bson:"roleRef"`
}

type RoleBinding struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sRoleBinding `bson:"obj" json:"obj"`
	KubeClusterId      string         `json:"kubeClusterId" bson:"kubeClusterId"`
}
