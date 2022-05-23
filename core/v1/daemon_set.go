package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sDaemonSet struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       DaemonSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     DaemonSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}
type DaemonSet struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sDaemonSet `bson:"obj" json:"obj"`
	AgentName          string       `bson:"agent_name" json:"agent_name"`
}
