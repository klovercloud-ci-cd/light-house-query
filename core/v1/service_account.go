package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sServiceAccount struct {
	TypeMeta                     `json:",inline" bson:",inline"`
	ObjectMeta                   `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Secrets                      []ObjectReference      `json:"secrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=secrets" bson:"secrets"`
	ImagePullSecrets             []LocalObjectReference `json:"imagePullSecrets,omitempty" protobuf:"bytes,3,rep,name=imagePullSecrets" bson:"imagePullSecrets"`
	AutomountServiceAccountToken *bool                  `json:"automountServiceAccountToken,omitempty" protobuf:"varint,4,opt,name=automountServiceAccountToken" bson:"automountServiceAccountToken"`
}

type ServiceAccount struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sServiceAccount `bson:"obj" json:"obj"`
	KubeClusterId      string            `json:"kubeClusterId" bson:"kubeClusterId"`
}
