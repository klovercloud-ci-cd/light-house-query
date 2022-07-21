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
	NamespaceCollection = "namespaceCollection"
)

type namespaceRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (n namespaceRepository) GetById(id, agent, processId string) v1.Namespace {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
			{"obj.metadata.uid": id},
		},
	}
	coll := n.manager.Db.Collection(NamespaceCollection)
	result := coll.FindOne(n.manager.Ctx, query, nil)
	elemValue := new(v1.Namespace)
	err := result.Decode(elemValue)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *elemValue
}

func (n namespaceRepository) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Namespace {
	var results []v1.Namespace
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := n.manager.Db.Collection(NamespaceCollection)
	findOptions := options.FindOptions{
		Sort: bson.M{"created_at": -1},
	}
	result, err := coll.Find(n.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Namespace)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (n namespaceRepository) GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Namespace, int64) {
	var results []v1.Namespace
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := n.manager.Db.Collection(NamespaceCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(n.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results, 0
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Namespace)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(n.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (n namespaceRepository) GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Namespace, int64) {
	var results []v1.Namespace
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.ownerReferences.uid": ownerReference},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := n.manager.Db.Collection(NamespaceCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(n.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results, 0
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Namespace)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(n.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

// NewNamespaceRepository returns repository.Namespace type repository
func NewNamespaceRepository(timeout int) repository.Namespace {
	return &namespaceRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
