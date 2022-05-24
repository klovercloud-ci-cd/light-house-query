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
	PVCollection = "pvCollection"
)

type persistentVolumeRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (p persistentVolumeRepository) Get(agent string, option v1.ResourceQueryOption) ([]v1.PersistentVolume, int64) {
	var results []v1.PersistentVolume
	query := bson.M{
		"$and": []bson.M{{"agent_name": agent}},
	}
	coll := p.manager.Db.Collection(PVCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(p.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.PersistentVolume)
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

// NewPersistentVolumeRepository returns repository.PersistentVolume type repository
func NewPersistentVolumeRepository(timeout int) repository.PersistentVolume {
	return &persistentVolumeRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
