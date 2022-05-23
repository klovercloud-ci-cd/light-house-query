package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sNetworkPolicy struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"Obj"`
	Spec       NetworkPolicySpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
}

type NetworkPolicy struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sNetworkPolicy `bson:"obj" json:"obj"`
	KubeClusterId      string           `json:"kubeClusterId" bson:"kubeClusterId"`
}
