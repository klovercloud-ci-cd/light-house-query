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

// PodCollection collection name
var (
	PodCollection = "podCollection"
)

type podRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (p podRepository) Get(agent string, option v1.ResourceQueryOption) ([]v1.Pod, int64) {
	var results []v1.Pod
	query := bson.M{
		"$and": []bson.M{{"agent_name": agent}},
	}
	coll := p.manager.Db.Collection(PodCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	result, err := coll.Find(p.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Pod)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(p.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

// NewPodRepository returns repository.Pod type repository
func NewPodRepository(timeout int) repository.Pod {
	return &podRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
