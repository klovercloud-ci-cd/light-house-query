package v1

import "github.com/go-bongo/bongo"

type Agent struct {
	bongo.DocumentBase `bson:",inline"`
	CompanyId          string `json:"company" bson:"company"`
	AgentName          string `json:"agent_name" bson:"agent_name"`
}
