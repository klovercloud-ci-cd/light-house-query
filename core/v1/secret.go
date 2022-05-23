package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sSecret struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Data       map[string][]byte `json:"data,omitempty" protobuf:"bytes,2,rep,name=data" bson:"data"`
	StringData map[string]string `json:"stringData,omitempty" protobuf:"bytes,4,rep,name=stringData" bson:"stringData"`
	Type       SecretType        `json:"type,omitempty" protobuf:"bytes,3,opt,name=type,casttype=SecretType" bson:"type"`
}

type Secret struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sSecret `bson:"obj" json:"obj"`
	AgentName          string    `bson:"agent_name" json:"agent_name"`
}
