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
	ServiceAccountCollection = "serviceAccountCollection"
)

type serviceAccountRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (s serviceAccountRepository) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.ServiceAccount, int64) {
	var results []v1.ServiceAccount
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.uid": ownerReference},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := s.manager.Db.Collection(ServiceAccountCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(s.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.ServiceAccount)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(s.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

// NewServiceAccountRepository returns repository.ServiceAccount type repository
func NewServiceAccountRepository(timeout int) repository.ServiceAccount {
	return &serviceAccountRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
