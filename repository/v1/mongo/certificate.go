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
	CertificateCollection = "certificateCollection"
)

type certificateRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (c certificateRepository) GetById(id, agent, processId string) v1.Certificate {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
			{"obj.metadata.uid": id},
		},
	}
	coll := c.manager.Db.Collection(CertificateCollection)
	result := coll.FindOne(c.manager.Ctx, query, nil)
	elemValue := new(v1.Certificate)
	err := result.Decode(elemValue)
	if err != nil {
		log.Println("[ERROR]", err)
		return *elemValue
	}
	return *elemValue
}

func (c certificateRepository) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.Certificate {
	var results []v1.Certificate
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := c.manager.Db.Collection(CertificateCollection)
	findOptions := options.FindOptions{
		Sort: bson.M{"created_at": -1},
	}
	result, err := coll.Find(c.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Certificate)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (c certificateRepository) GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.Certificate, int64) {
	var results []v1.Certificate
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := c.manager.Db.Collection(CertificateCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(c.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results, 0
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Certificate)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
		return results, 0
	}
	return results, count
}

func (c certificateRepository) GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.Certificate, int64) {
	var results []v1.Certificate
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.ownerReferences.uid": ownerReference},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := c.manager.Db.Collection(CertificateCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(c.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results, 0
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Certificate)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
		return results, 0
	}
	return results, count
}

// NewCertificateRepository returns repository.Certificate type repository
func NewCertificateRepository(timeout int) repository.Certificate {
	return &certificateRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
