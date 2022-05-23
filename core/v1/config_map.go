package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sConfigMap struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Data       map[string]string `json:"data,omitempty" protobuf:"bytes,2,rep,name=data" bson:"data"`
	BinaryData map[string][]byte `json:"binaryData,omitempty" protobuf:"bytes,3,rep,name=binaryData" bson:"binaryData"`
}

type ConfigMap struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sConfigMap `bson:"obj" json:"obj"`
	KubeClusterId      string       `json:"kubeClusterId" bson:"kubeClusterId"`
}
