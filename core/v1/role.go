package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sRole struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata" `
	Rules      []PolicyRule `json:"rules" protobuf:"bytes,2,rep,name=rules" bson:"rules"`
}

type Role struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sRole `bson:"obj" json:"obj"`
	AgentName          string  `bson:"agent_name" json:"agent_name"`
}
