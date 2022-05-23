package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sNode struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       NodeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     NodeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}
type Node struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sNode `bson:"obj" json:"obj"`
	KubeClusterId      string  `json:"kubeClusterId" bson:"kubeClusterId"`
}
