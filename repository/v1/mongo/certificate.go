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

func (c certificateRepository) Get(agent string, option v1.ResourceQueryOption) ([]v1.Certificate, int64) {
	var results []v1.Certificate
	query := bson.M{
		"$and": []bson.M{{"agent_name": agent}},
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
	}
	return results, count
}

func (c certificateRepository) GetByOwnerReference(agent, ownerReference string, option v1.ResourceQueryOption) ([]v1.Certificate, int64) {
	var results []v1.Certificate
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.uid": ownerReference},
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
