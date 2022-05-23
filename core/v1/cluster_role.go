package v1

import (
	"github.com/go-bongo/bongo"
)

type K8sClusterRole struct {
	TypeMeta        `json:",inline" bson:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Rules           []PolicyRule     `json:"rules" protobuf:"bytes,2,rep,name=rules" bson:"rules"`
	AggregationRule *AggregationRule `json:"aggregationRule,omitempty" protobuf:"bytes,3,opt,name=aggregationRule" bson:"aggregationRule"`
}
type ClusterRole struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sClusterRole `bson:"obj" json:"obj"`
	AgentName          string         `bson:"agent_name" json:"agent_name"`
}
