package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sDeployment struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       DeploymentSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     DeploymentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}

type Deployment struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sDeployment `bson:"obj" json:"obj"`
	KubeClusterId      string        `json:"kubeClusterId" bson:"kubeClusterId"`
}
