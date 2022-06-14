package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

var (
	AgentIndexCollection = "agentIndexCollection"
)

type agentRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (a agentRepository) Get(companyId string) []v1.Agent {
	var results []v1.Agent
	query := bson.M{
		"$and": []bson.M{
			{"company": companyId},
		},
	}
	coll := a.manager.Db.Collection(AgentIndexCollection)
	result, err := coll.Find(a.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Agent)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

// NewAgentRepository returns repository.Agent type repository
func NewAgentRepository(timeout int) repository.Agent {
	return &agentRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
