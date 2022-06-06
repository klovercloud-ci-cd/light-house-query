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
	IngressCollection = "ingressCollection"
)

type ingressRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (i ingressRepository) GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Ingress, int64) {
	var results []v1.Ingress
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := i.manager.Db.Collection(IngressCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(i.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Ingress)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(i.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (i ingressRepository) GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Ingress, int64) {
	var results []v1.Ingress
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.uid": ownerReference},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := i.manager.Db.Collection(IngressCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(i.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Ingress)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(i.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

// NewIngressRepository returns repository.Ingress type repository
func NewIngressRepository(timeout int) repository.Ingress {
	return &ingressRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
