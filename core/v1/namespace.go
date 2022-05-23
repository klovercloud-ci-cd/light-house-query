package v1

import (
	"github.com/go-bongo/bongo"
)

type FinalizerName string
type NamespacePhase string

const (
	FinalizerKubernetes  FinalizerName  = "kubernetes"
	NamespaceActive      NamespacePhase = "Active"
	NamespaceTerminating NamespacePhase = "Terminating"
)

type K8sNamespace struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       NamespaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     NamespaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}

type Namespace struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sNamespace `bson:"obj" json:"obj"`
	AgentName          string       `bson:"agent_name" json:"agent_name"`
}
