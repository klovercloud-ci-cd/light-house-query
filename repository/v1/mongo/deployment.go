package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	DeploymentCollection = "deploymentCollection"
)

type deploymentRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (d deploymentRepository) CountDeploymentsByCompanyIdAndGroupByAgent(companyId string) map[string]int64 {
	results := make(map[string]int64)
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.labels.company": companyId},
		},
	}
	coll := d.manager.Db.Collection(DeploymentCollection)
	result, err := coll.Find(d.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.DeploymentShortDto)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if val, ok := results[elemValue.AgentName]; ok {
			results[elemValue.AgentName] = val + 1
		} else {
			results[elemValue.AgentName] = 1
		}
	}
	return results
}

func (d deploymentRepository) GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Deployment, int64) {
	var results []v1.Deployment
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := d.manager.Db.Collection(DeploymentCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(d.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Deployment)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(d.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (d deploymentRepository) GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Deployment, int64) {
	var results []v1.Deployment
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.uid": ownerReference},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := d.manager.Db.Collection(DeploymentCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(d.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Deployment)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(d.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

// NewDeploymentRepository returns repository.Deployment type repository
func NewDeploymentRepository(timeout int) repository.Deployment {
	return &deploymentRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
