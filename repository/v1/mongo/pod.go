package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
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

func GetPagination(size int64, page int64, limit int64) (int64, int64) {
	startIndex := page * limit
	if startIndex >= size {
		return 0, 0
	}
	if size <= startIndex+limit {
		return startIndex, size
	}
	return startIndex, startIndex + limit
}

func (p podRepository) Get(agent string) ([]v1.Pod, int64) {
	var results []v1.Pod
	query := bson.M{
		"$and": []bson.M{{"agent_name": agent}},
	}
	coll := p.manager.Db.Collection(PodCollection)
	result, err := coll.Find(p.manager.Ctx, query, nil)
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
	return results, int64(len(results))
}

// NewPodRepository returns repository.Pod type repository
func NewPodRepository(timeout int) repository.Pod {
	return &podRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
