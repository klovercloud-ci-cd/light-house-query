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
	ReplicaSetCollection = "replicaSetCollection"
)

type replicaSetRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (r replicaSetRepository) Get(agent string, option v1.ResourceQueryOption) ([]v1.ReplicaSet, int64) {
	var results []v1.ReplicaSet
	query := bson.M{
		"$and": []bson.M{{"agent_name": agent}},
	}
	coll := r.manager.Db.Collection(ReplicaSetCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(r.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.ReplicaSet)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(r.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

// NewReplicaSetRepository returns repository.ReplicaSetRepository type repository
func NewReplicaSetRepository(timeout int) repository.ReplicaSet {
	return &replicaSetRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}